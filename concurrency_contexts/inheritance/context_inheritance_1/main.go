package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создается контекст, который завершится через 2 секунды.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// вызывается функция с таким контекстом
	makeRequest(ctx)
}

func makeRequest(ctx context.Context) {
	// создаем таймер на 5 секунд
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	// создаем новый контекст, который завершится через 10 секунд. 
	// Причем этот новый контекст отнаследован от контекста ctx (созданного в main)
	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// здесь будет напечатано canceled, поскольку ctx отработает быстрее, 
	// чем таймер, и по цепочке отменится и дочерний контекст. 
	select {
	case <-newCtx.Done():
		fmt.Println("canceled")
	case <-timer.C:
		fmt.Println("timer")
	}
}
