package main

import (
	"context"
	"fmt"
)

func main() {
	traceCtx := context.WithValue(context.Background(), "trace_id", "12-21-33")
	makeRequest(traceCtx)
}

func makeRequest(ctx context.Context) {
	oldValue, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Println(oldValue) // 12-21-33
	}

	// отнаследуем новый контекст от ctx, созданного в main
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Будет напечатано 12-21-33
	newValue, ok := newCtx.Value("trace_id").(string)
	if ok {
		fmt.Println(newValue)
	}
}
