package main

import (
	"context"
	"errors"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// Задача на оптимизацию cache-а - одна из популярных задач на собеседованиях. 
// Задача заключается в том, что необходимо написать обертку над базой данных. 
// При реализации этой задачи, необходимо задавать вопросы об ограничениях: 
// 1) Размер оперативной памяти 
// 2) Как много данных в базе 
// 3) Как должна производится инвалидация 
// 4) Какие типы данных 
// 5) Сколько мы можем позволить, чтобы данные устаревали 
// 6) Будут ли какие механизмы вытеснения 
// 7) Будем ли мы ограничивать кэш по размеру 

// Здесь рассматривается достаточно простой вариант: на стороне кэша памяти очень 
// много, а сторонняя база данных достаточно маленькая. Причем у нас нет никаких 
// дополнительных расширений, нам просто необходимо раз в некоторый временной 
// промежуток обновлять ключи. Это можно реализовать в нескольких вариантах: 
// 1) Хранить некую временную метку для объекта и раз, например, в две минуты ходить и 
// обновлять данные. 
// 2) Обновлять данные раз в некоторый временной промежуток. 

// Здесь представлен код, который повзоляет в асинронном режиме производить инвалидацию 
// кэша раз в минуту. 

const retriesNumber = 3
const retriesInitialPauseDuration = time.Millisecond * 100
const requestTimeout = time.Second * 5
const syncInterval = time.Minute

// Представим, что есть некоторый интерфейс, который реализует взаимодействие с базой данных.
type RedisDatabase interface {
	// Метод,  который по ключу возвращает некоторое значение.
	Get(context.Context, string) (string, error)
	// Метод, который позволяет получить множество значений по множеству ключей. 
	MGet(context.Context, []string) ([]*string, error)
	// Метод, который возвращает множество ключей. 
	Keys(context.Context) ([]string, error)
}

type RedisDatabaseWithCache struct {
	database RedisDatabase
	group    singleflight.Group

	mutex sync.RWMutex
	cache map[string]string
}

func NewRedisDatabaseWithCache(ctx context.Context, database RedisDatabase) (*RedisDatabaseWithCache, error) {
	// Проверяется, что контекст уже отменен.
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// Если database некорректный, то тоже вернем ошибку. 
	if database == nil {
		return nil, errors.New("incorrect database")
	}

	// Если все корректно, то создаем структуру, которая будет описывать кэш, 
	// куда добавляется database и мапа с key/value в виде string. 
	c := &RedisDatabaseWithCache{
		database: database,
		cache:    make(map[string]string),
	}

	// В фоне запускается горутина, которая будет синхронизировать данные. 
	go c.synchronize(ctx)
	return c, nil
}

func (c *RedisDatabaseWithCache) synchronize(ctx context.Context) {
	// Запускается таймер, по которому раз в минуту будет происходить тик и синхронизация. 
	ticker := time.NewTicker(syncInterval)
	defer ticker.Stop()

	// То есть в бесконечном цикле либо проверяется, что закончится контекст, 
	// либо произойдет тик и надо будет синхронизироваться 
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Тут можно будет задать вопрос: а что делать, 
			// если синхронизация данных будет длится дольше 
			// минуты и тики будут накапливаться. 
			// Одна из реализаций - сделать так, чтобы тамауты 
			// были бы меньше, чем минута. 
			c.synchronizeImpl(ctx)
		}
	}
}

func (c *RedisDatabaseWithCache) synchronizeImpl(ctx context.Context) {
	var keys []string
	// Здесь вызывается некоторая обертка, которая реализует механизм retry
	err := withRetries(ctx, func(ctx context.Context) error {
		var err error
		keys, err = c.database.Keys(ctx)
		return err
	})

	// Если вернулась ошибка - никакой из retry не пробился, фиксируем эти данные
	if err != nil {
		return // logs, metrics
	}

	// Ананлогично по всем ключам с оберткой для реализации retry получаем 
	// множество значений. 
	var values []*string
	err = withRetries(ctx, func(ctx context.Context) error {
		var err error
		values, err = c.database.MGet(ctx, keys)
		return err
	})

	if err != nil {
		return // logs, metrics
	}

	// Очень важно не добавлять новые элементы в мапу под блокировками, 
	// поскольку это будет работать гораздо медленнее. Дргуие горутины, 
	// которые будут обращаться к кэшу будут ждать. 
	// Поэтому здесь создается временная мапа, в которую записываются ключи 
	// и значения. 
	cache := make(map[string]string, len(keys))
	for idx, key := range keys {
		value := values[idx]
		if value != nil {
			cache[key] = *value
		}
	}

	// Далее используется мелкогранулярная блокировка, чтобы заменить старую 
	// мапу на новую мапу.  
	c.mutex.Lock()
	c.cache = cache
	c.mutex.Unlock()
}

// Get need to proxy only one method from interface
func (c *RedisDatabaseWithCache) Get(ctx context.Context, key string) (string, error) {
	var found bool
	var value string
	// withLock - тоже некоторая обертка, которая сначала назначает мьютекс, 
	// затем выполняется функцию, а затем снимает мьютекс. 
	withLock(c.mutex.RLocker(), func() {
		value, found = c.cache[key]
	})

	// Если значение найдено в кэше - возвращаем его. 
	if found {
		return value, nil
	}

	// Данная конструкция реализует singleFlight - то есть, когда одна горутина 
	// достает данные, а остальные ждут ее исполнения и забирают значение. 
	valueFromDB, err, _ := c.group.Do(key, func() (any, error) {
		var value string
		// Идем в базу данных, чтобы достать значение оттуда. 
		err := withRetries(ctx, func(ctx context.Context) error {
			var err error
			value, err = c.database.Get(ctx, key)
			return err
		})

		if err != nil {
			return "", err
		}

		// Если значение появилось - возвращаем его пользователю и записываем его в кэш
		withLock(&c.mutex, func() {
			c.cache[key] = value
		})

		return value, err
	})

	return valueFromDB.(string), err
}

func withLock(mutex sync.Locker, action func()) {
	if action == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	action()
}

// withRetries - функция, которая принимает контекст и некоторую функцию
func withRetries(ctx context.Context, action func(ctx context.Context) error) error {
	if action == nil {
		return errors.New("incorrect action")
	}

	var err error
	// Идем циклом по количеству retry (в данном случае это 3)
	for idx := 1; idx <= retriesNumber; idx++ {
		// Запускаем таймаут на 1 запрос 
		ctx, cancel := context.WithTimeout(ctx, requestTimeout)
		defer cancel()

		// Вызываем функцию с контекстом 
		err := action(ctx)
		if err == nil {
			return nil
		}
		// После retry выполняется Back off - то есть откладывается момент следующего retry 
		time.Sleep(time.Duration(idx) * retriesInitialPauseDuration)
	}

	return err
}
