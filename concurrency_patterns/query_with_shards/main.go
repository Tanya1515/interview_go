package main

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Здесь представлен код, который позволяет сделать распределенный 
// запрос в несколько шардов или в несколько сервисов6 в которых 
// находятся абсолютно разная информация. 

// Здесь стоит учитывать, что, когда выполняется cross-шардовый запрос, 
// то если один из запросов не выполнился - ждать остальные бессмысленно, 
// поскольку необходимо собрать данные из всех шардов. 
// А если один из запросов не выполнился, то результат будет неполным. 

type Database interface {
	Query(query string) (string, error) // simple interface for example
}

// DistributedQuery принимает список шардов, запрос и возвращает список ответов и ошибку. 
func DistributedQuery(shards []Database, query string) ([]string, error) {
	var mutex sync.Mutex
	// Создается результирующий срез под каждый ответ среди всех шардов. 
	responses := make([]string, 0, len(shards))
	// Создаем errorgroup context. 
	group, ctx := errgroup.WithContext(context.TODO()) // nice to have parent context

	// Итерируемся по каждому из шардов, создавая на каждой итерации новую горутину. 
	for _, shard := range shards {
		group.Go(func() error {
			// Создается структура result 
			type result struct {
				response string
				err      error
			}

			// Создается канал результатов буферезированный!!!
			resultCh := make(chan result, 1)
			// В отдельной горутине запускается запрос к шарду, а в результирующий 
			// канал записывается объект типа result. 
			go func() {
				response, err := shard.Query(query)
				resultCh <- result{response: response, err: err}
			}()

			// Далее в select-е ожидаем либо завершение контекста, 
			// либо приходят данные в результирующий канал 
			select {
			// Выполнение этого контекста означает выполнение контекста 
			// для errorgroup (групповой контекст), 
			// то есть одна из горутин завершила работу с ошибкой.

			// Важно отметить, что канал здесь сделан буферезированным, 
			// поскольку если горутина завершит свою работу по контексту, 
			// а горутина с запросом не успеет отработать до этого 
			// момента, то будет утечка памяти. 
			case <-ctx.Done():
				return ctx.Err()
			// Если же пришло значение в канал, то проверяем ошибку, а затем 
			// под mutex-ом записываес значение в responses (поскольку это 
			// переменная, которая шарится на несколько горутин). 
			case result := <-resultCh:
				if result.err != nil {
					return result.err
				}

				mutex.Lock()
				responses = append(responses, result.response)
				mutex.Unlock()

				return nil
			}
		})
	}

	// Если в одной из горутин произошла ошибка - возвращаем ее, 
	// в ином варианте - возвращаем слайс результатов
	if err := group.Wait(); err != nil {
		return nil, err
	} else {
		return responses, nil
	}
}

func main() {
	shards := []*ClickHouseDatabase{
		NewClickHouseDatabase("127.0.0.1:5432"),
		NewClickHouseDatabase("127.0.0.2:5432"),
		NewClickHouseDatabase("127.0.0.3:5432"),
	}

	response, err := DistributedQuery(shards, "query to clickhouse...")
	_ = response
	_ = err
}
