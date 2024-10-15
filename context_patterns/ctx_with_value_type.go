package context_patterns

import (
	"context"
	"fmt"
)

func main() {
	{
		ctx := context.WithValue(context.Background(), "key", "1")
		ctx = context.WithValue(ctx, "key", "2")

		fmt.Printf("value: %s \n", ctx.Value("key").(string))
	}

	type key1 string
	type key2 string

	var k1 key1 = "key"
	var k2 key2 = "key"

	ctx := context.WithValue(context.Background(), k1, "1")
	ctx = context.WithValue(ctx, k2, "2")

	fmt.Printf("value: %s \n", ctx.Value(k1).(string))
	fmt.Printf("value: %s \n", ctx.Value(k2).(string))
}
