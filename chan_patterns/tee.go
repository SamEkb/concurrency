package chan_patterns

func TeeSplit(inputChan chan int, number int) []chan int {
	outputChan := make([]chan int, number)

	for i := 0; i < number; i++ {
		outputChan[i] = make(chan int)
	}

	go func() {
		for value := range inputChan {
			for _, ch := range outputChan {
				ch <- value
			}
		}

		for _, ch := range outputChan {
			close(ch)
		}
	}()

	return outputChan
}

//func main() {
//	ch := make(chan int)
//
//	go func() {
//		defer close(ch)
//
//		for i := 0; i < 3; i++ {
//			ch <- i
//		}
//	}()
//
//	result := TeeSplit(ch, 2)
//	wg := sync.WaitGroup{}
//	wg.Add(len(result))
//
//	go func() {
//		defer wg.Done()
//		for value := range result[0] {
//			fmt.Println("ch_1: ", value)
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		for value := range result[1] {
//			fmt.Println("ch_2: ", value)
//		}
//	}()
//
//	wg.Wait()
//}
