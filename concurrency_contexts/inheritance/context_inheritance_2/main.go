package main

import (
	"context"
	"fmt"
	"time"
)

// Контексты работают как матрешка: если контекст отменяется, то это влияет только на его детей. 

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, cancel = context.WithCancel(ctx)
	cancel()

	if ctx.Err() != nil {
		// canceled не будет напечатан, поскольку, если отменяется 
		// дочерний контекст, то на родительский это никак не будет влиять. 
		fmt.Println("canceled") 
	}
}
