package chan_patterns

import (
	"sync"
)

func MergeChannels(channels ...chan int) <-chan int {
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	result := make(chan int)
	for _, channel := range channels {
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				result <- value
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

//func main() {
//	chan1 := make(chan int)
//	chan2 := make(chan int)
//	chan3 := make(chan int)
//
//	go func() {
//		defer func() {
//			close(chan1)
//			close(chan2)
//			close(chan3)
//		}()
//
//		for i := 0; i < 100; i += 3 {
//			chan1 <- i
//			chan2 <- i + 1
//			chan3 <- i + 2
//		}
//	}()
//
//	for value := range MergeChannels(chan1, chan2, chan3) {
//		fmt.Println(value)
//	}
//}
