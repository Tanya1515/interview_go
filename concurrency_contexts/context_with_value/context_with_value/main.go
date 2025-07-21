package main

import (
	"context"
	"fmt"
)

func main() {
	traceCtx := context.WithValue(context.Background(), "trace_id", "12-21-33")
	makeRequest(traceCtx)

	// добавление trace_id у дочернего контекста с новым значением 
	// никак не поменяет значение trace_id у родительского контекста. 
	oldValue, ok := traceCtx.Value("trace_id").(string)
	if ok {
		fmt.Println("mainValue", oldValue) // 12-21-33
	}
}

func makeRequest(ctx context.Context) {
	oldValue, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Println("oldValue", oldValue) // 12-22-33
	}

	// создаем дочерний контекст со значением trace_id равным 22-22-22. 
	// То есть старое значение было перекрыто.  
	newCtx := context.WithValue(ctx, "trace_id", "22-22-22")
	newValue, ok := newCtx.Value("trace_id").(string)
	if ok {
		fmt.Println("newValue", newValue) // 22-22-22
	}
}
