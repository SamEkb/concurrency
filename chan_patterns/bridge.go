package chan_patterns

import (
	"sync"
)

func Bridge(inputChan chan chan string) <-chan string {
	outChan := make(chan string)
	wg := sync.WaitGroup{}

	go func() {
		for channel := range inputChan {
			wg.Add(1)
			go func(ch chan string) {
				defer wg.Done()
				for value := range ch {
					outChan <- value
				}
			}(channel)
		}
		wg.Wait()
		close(outChan)
	}()

	return outChan
}

//func main() {
//	channels := make(chan chan string)
//	go func() {
//		firstChan := make(chan string, 5)
//		for i := 0; i < 5; i++ {
//			firstChan <- fmt.Sprintf("channel one: %d", i)
//		}
//		close(firstChan)
//
//		secondChan := make(chan string, 5)
//		for i := 0; i < 5; i++ {
//			secondChan <- fmt.Sprintf("channel two: %d", i)
//		}
//		close(secondChan)
//
//		channels <- firstChan
//		channels <- secondChan
//		close(channels)
//	}()
//
//	for value := range bridge(channels) {
//		fmt.Println(value)
//	}
//}
