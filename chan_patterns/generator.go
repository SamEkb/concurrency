package chan_patterns

func GenerateWithClosure(number int) func() int {
	return func() int {
		r := number
		number++
		return r
	}
}

func GenerateWithChan(start, end int) <-chan int {
	outChan := make(chan int)

	go func() {
		defer close(outChan)
		for i := start; i <= end; i++ {
			outChan <- i
		}
	}()

	return outChan
}

//func main() {
//	generator := GenerateWithClosure(0)
//	for i := 0; i <= 200; i++ {
//		fmt.Println(generator())
//	}
//
//	for number := range GenerateWithChan(0, 10) {
//		fmt.Println(number)
//	}
//}
