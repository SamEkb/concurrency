package chan_patterns

func filter[T any](inputChan <-chan T, predicate func(T) bool) <-chan T {
	outputChan := make(chan T)

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
