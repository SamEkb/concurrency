package context_patterns

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", "1")
	makeReq(ctx)

	// value не изменится
	oldValue, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Println(oldValue)
	}
}

func makeReq(ctx context.Context) {
	oldValue, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Println(oldValue)
	}

	//создаем новый контекст на основе старого и добавляем такойже ключ
	newCtx := context.WithValue(ctx, "trace_id", "2")
	newValue, ok := newCtx.Value("trace_id").(string)
	if ok {
		fmt.Println("newValue", newValue)
	}
}
