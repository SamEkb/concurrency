package context_patterns

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type DistributedDatabase struct {
}

// Query имитируем обращение в бд
func (d *DistributedDatabase) Query(address, key string) string {
	return fmt.Sprintf("finished with address: [%s]", address)
}

var database DistributedDatabase

func Query(ctx context.Context, result chan string, address string) {
	randomTime := time.Duration(rand.Intn(5000)) * time.Millisecond

	timer := time.NewTimer(randomTime)
	select {
	case <-timer.C:
		result <- database.Query(address, "")
	case <-ctx.Done():

		fmt.Printf("cancelled with address: [%s] \n", address)
	}
}

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	wg := sync.WaitGroup{}
//
//	addresses := []string{
//		"127.0.0.1",
//		"127.0.0.2",
//		"127.0.0.3",
//		"127.0.0.4",
//		"127.0.0.5",
//		"127.0.0.6",
//		"127.0.0.7",
//		"127.0.0.8",
//		"127.0.0.9",
//	}
//
//	result := make(chan string)
//	for _, address := range addresses {
//		wg.Add(1)
//		go func(address string) {
//			defer wg.Done()
//			Query(ctx, result, address)
//		}(address)
//	}
//
//	fmt.Println(<-result)
//
//	cancel()
//	wg.Wait()
//}
