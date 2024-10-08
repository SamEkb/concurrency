package chan_patterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ShardedDatabase struct {
}

func (s *ShardedDatabase) Query(shard, key string) string {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	return fmt.Sprintf("[%s]: value", shard)
}

var shardedDatabase ShardedDatabase

func Get(shards []string, query string) []string {
	resultCh := make(chan string, len(shards))
	wg := sync.WaitGroup{}

	for _, shard := range shards {
		wg.Add(1)
		go func(shard string) {
			defer wg.Done()
			res := shardedDatabase.Query(shard, query)
			resultCh <- res
		}(shard)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []string
	for result := range resultCh {
		results = append(results, result)
	}

	return results
}

//func main() {
//	addresses := []string{
//		"shard_1",
//		"shard_2",
//		"shard_3",
//	}
//
//	value := Get(addresses, "")
//	fmt.Println(value)
//}
