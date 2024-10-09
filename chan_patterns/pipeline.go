package chan_patterns

func generate(numbers ...int) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)
		for _, num := range numbers {
			outputChan <- num
		}
	}()

	return outputChan
}

func multiply(inputChan <-chan int, factor int) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)
		for num := range inputChan {
			outputChan <- num * factor
		}
	}()

	return outputChan
}

//func main() {
//	for num := range multiply(generate(1, 2, 3, 4, 5), 2) {
//		fmt.Println(num)
//	}
//}
