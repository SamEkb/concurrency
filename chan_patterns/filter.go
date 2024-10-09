package chan_patterns

func filter(inputChan <-chan int, predicate func(int) bool) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)
		for value := range inputChan {
			if predicate(value) {
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
//	for value := range filter(ch, func(i int) bool {
//		return i%2 == 0 && i > 0
//	}) {
//		fmt.Println(value)
//	}
//}
