package main

import (
	"context"
	"fmt"
)

// Результат работы программы: 

// string = value2 
// key1 = value1 
// key2 = value2


func main() {
	{
		ctx := context.WithValue(context.Background(), "key", "value1")
		ctx = context.WithValue(ctx, "key", "value2")

		// здесь дочерний контекст перекроет значение родительского контекста key 
		fmt.Println("string =", ctx.Value("key").(string))
	}
	{
		type key1 string // type definition, not type alias (type definition - это отдельный тип)
		type key2 string // type definition, not type alias
		// создается несколько констант типа key1 и key2
		const k1 key1 = "key" 
		const k2 key2 = "key"

		// далее создаются два контекста, у которых будут value типа key1 и key2 соответственно
		ctx := context.WithValue(context.Background(), k1, "value1")
		ctx = context.WithValue(ctx, k2, "value2")

		// здесь же, в силу того, что типы у значений контекста разные - то и значения будут разные
		// Но если в качестве key1 и key2 использовать type alias-ы:
		// type key1 = string 
		// type key2 = string, то значения перекроют друг друга и будет напечатано: 
		// key1 = value2 
		// key2 = value2

		// Такая проблема былвает выползает, если в контекст пробрасывается некоторое значение, 
		// а затем разработчик подключает библиотеку, которая перезатирает это значение. 
		fmt.Println("key1 =", ctx.Value(k1).(string))
		fmt.Println("key2 =", ctx.Value(k2).(string))
	}
}
