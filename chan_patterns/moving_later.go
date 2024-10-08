package chan_patterns

import (
	"fmt"
	"math/rand"
	"time"
)

type DistributedDatabase struct {
}

// Query имитируем обращение в бд
func (d *DistributedDatabase) Query(address, key string) string {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	return fmt.Sprintf("[%s]: value", address)
}

var database DistributedDatabase

// Query Полуаем данные из одной реплики, остальные не ждем
func Query(addresses []string, query string) string {
	result := make(chan string)

	for _, address := range addresses {
		go func(address string) {
			select {
			case result <- database.Query(address, query):
			default:
				return
			}
		}(address)
	}

	return <-result
}

//func main() {
//	addresses := []string{
//		"127.0.0.1",
//		"127.0.0.2",
//		"127.0.0.3",
//	}
//
//	value := Query(addresses, "")
//	fmt.Println(value)
//}
