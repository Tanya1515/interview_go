package main

import (
	"errors"
	"sync"
)

// Паттерн Worker Pool - основная идея заключается в том, что запускается 
// некоторое множество горутин, которые ждут исполнения задач. Когда 
// горутина выполнила некоторую задачу, горутина не завершается, а продолжает 
// ждать, когда ей делегируют другую задачу. На этапе решения задачи на 
// собеседовании нужно обговорить детали, на базе которых будет реализована задача: 

// 1) Когда добавляется задача, а готовых воркеров либо нет, либо они все заняты, 
// как должна вести себя программа: блокироваться или возвращать ошибку. 
// 2) Как завершить воркер: при помощи контекста или при помощи контекста или 
// при помощи shutdown - то есть дождаться, когда все воркеры завершаться. 
// 3) Как быть с задачами, которые находятся в буфере? 


type WorkerPool struct {
	tasksCh chan func()
	mutex   sync.RWMutex

	closed      bool
	closeDoneCh chan struct{}
}

// Констуктор worker pool-а
func NewWorkerPool(workersNumber int) (*WorkerPool, error) {
	if workersNumber <= 0 {
		return nil, errors.New("incorrect workers number")
	}

	// Здесь создается канал закрытия и создается буферезированный канал, 
	// куда будут складываться задачи. 
	wp := &WorkerPool{
		closeDoneCh: make(chan struct{}),
		tasksCh:     make(chan func(), workersNumber),
	}

	// Запускается асинхронная горутина, которая начинает обрабатывать задачи.
	go wp.processTasks(workersNumber)
	return wp, nil
}

// Функция которая запускает горутины, которые будут на фоне 
// выполнять пришедшие задачи
func (wp *WorkerPool) processTasks(workersNumber int) {
	var wg sync.WaitGroup
	// Создаем WaitGroup, в которой число равно workerNumber
	wg.Add(workersNumber)

	// В цикле создаем горутины, которые будут читать из канала задачи, 
	// пока канал задач не будет закрыт.  
	for i := 0; i < workersNumber; i++ {
		go func() {
			defer wg.Done()
			for task := range wp.tasksCh {
				task()
			}
		}()
	}

	// Здесь ожидают, пока все горутины закончат работу 
	wg.Wait()
	// Закрывается канал закрытия 
	close(wp.closeDoneCh)
}

// AddTask add task to pool
// Функция, которая добавляет задачу в очередь.
func (wp *WorkerPool) AddTask(task func()) error {
	// Здесь проверяется является ли задача валидной. 
	if task == nil {
		return errors.New("incorrect task")
	}


	// Здесь используется мьютекс на чтение, то есть разрешается множеству 
	// горутин заходить и писать значения в канал, то есть удерживаем 
	// блокировку на чтение и запрещаем писать. Поскольку, как только 
	// в методе Close поменялся флаг wg.closed - горутина не сможет 
	// записать в канал. 
	wp.mutex.RLock()
	defer wp.mutex.RUnlock()

	if wp.closed {
		return errors.New("pool is closed")
	}

	// То есть мьютекс защищает от ситации, когда горутина прервалась в этом месте,
	// а канал закрылся в методе Closed

	// Здесь производится попытка записи задачи в канал, если все 
	// в порядке, то возвращаем nil, если есть какая-то ошибка - 
	// возвращаем pool is full.  
	select {
	case wp.tasksCh <- task:
		return nil
	default:
		return errors.New("pool is full")
	}
}

// Close close pool and wait all the tasks
// Данный метод будет ждать, пока все worker-ы остановятся, более того необходимо 
// будет дождаться, пока все задачи из буфера тоже разгребуться. 
func (wp *WorkerPool) Close() {
	// Здесь проверяется, а закрыт ли worker pool, 
	// чтобы не было паники о закрытии уже закрытого канала. 
	// При этом обращение к структуре производится без 
	// блокировки - поскольку это операция чтения. 
	if wp.closed {
		return
	}

	// Далее выставляем флаг, что канал закрыт, используя при этом мьютекс. 
	// Этот флаг используется для того, чтобы никакая из горутин не могла 
	// записать в закрытый канал.
	wp.mutex.Lock()
	wp.closed = true
	wp.mutex.Unlock()

	close(wp.tasksCh)
	// Здесь метод Close ожидает closeDoneChannel, то есть когда все воркеры доработают
	<-wp.closeDoneCh
}
