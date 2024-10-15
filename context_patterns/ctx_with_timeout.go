package context_patterns

import (
	"context"
	"fmt"
	"time"
)

func makeRequest(ctx context.Context) {
	timer := time.NewTimer(time.Second * 2)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("cancel with timer")
	case <-ctx.Done():
		fmt.Println("cancel with context")
	}
}

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	makeRequest(ctx)
//}
