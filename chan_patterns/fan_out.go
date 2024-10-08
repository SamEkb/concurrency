package chan_patterns

func SplitChannels(inputChan chan int, number int) []chan int {

	outputChan := make([]chan int, number)

	for i := 0; i < number; i++ {
		outputChan[i] = make(chan int)
	}

	go func() {
		idx := 0
		for value := range inputChan {
			outputChan[idx] <- value
			idx = (idx + 1) % number
		}

		for _, ch := range outputChan {
			close(ch)
		}
	}()

	return outputChan
}

//func main() {
//	chan1 := make(chan int)
//
//	go func() {
//		defer close(chan1)
//
//		for i := 0; i < 10; i++ {
//			chan1 <- i
//		}
//	}()
//
//	result := SplitChannels(chan1, 2)
//
//	wg := sync.WaitGroup{}
//	wg.Add(len(result))
//
//	go func() {
//		defer wg.Done()
//		for value := range result[0] {
//			fmt.Println("ch1: ", value)
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		for value := range result[1] {
//			fmt.Println("ch2: ", value)
//		}
//	}()
//
//	wg.Wait()
//}
