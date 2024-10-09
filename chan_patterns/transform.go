package chan_patterns

func transform(inputChan <-chan int) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)
		for value := range inputChan {
			outputChan <- value * 10
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
//	for value := range transform(ch) {
//		fmt.Println(value)
//	}
//
//}
