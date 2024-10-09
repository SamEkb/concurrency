package chan_patterns

import (
	"fmt"
	"sync"
)

func parse(inputChan <-chan string) <-chan string {
	outputChan := make(chan string)

	go func() {
		defer close(outputChan)
		for data := range inputChan {
			outputChan <- fmt.Sprintf("parsed %s", data)
		}
	}()

	return outputChan
}

func send(inputChan <-chan string, number int) <-chan string {
	outputChan := make(chan string)
	channels := split(inputChan, number)

	wg := sync.WaitGroup{}
	wg.Add(number)

	for idx, channel := range channels {
		go func(ch <-chan string, idx int) {
			defer wg.Done()
			for data := range ch {
				outputChan <- fmt.Sprintf("sent data: %s to replica %d", data, idx)
			}
		}(channel, idx)
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	return outputChan
}

func split(inputChan <-chan string, number int) []chan string {
	outputChan := make([]chan string, number)

	for i := 0; i < number; i++ {
		outputChan[i] = make(chan string)
	}

	go func() {
		idx := 0
		for data := range inputChan {
			outputChan[idx] <- data
			idx = (idx + 1) % number
		}

		for _, ch := range outputChan {
			close(ch)
		}
	}()

	return outputChan
}

//func main() {
//	ch := make(chan string)
//
//	go func() {
//		defer close(ch)
//		for i := 1; i <= 10; i++ {
//			ch <- fmt.Sprintf("address %d", i)
//		}
//	}()
//
//	for result := range send(parse(ch), 2) {
//		fmt.Println(result)
//	}
//}
