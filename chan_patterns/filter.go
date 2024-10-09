package chan_patterns

func filter(inputChan <-chan int) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)
		for value := range inputChan {
			if value%2 == 0 && value > 0 {
				outputChan <- value
			}
		}
	}()

	return outputChan
}

//func main() {
//	ch := make(chan int)
//
//	go func() {
//		defer close(ch)
//		for i := 0; i < 10; i++ {
//			ch <- i
//		}
//	}()
//
//	for value := range filter(ch) {
//		fmt.Println(value)
//	}
//}
