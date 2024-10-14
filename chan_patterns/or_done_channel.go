package chan_patterns

func OrDone(inChan chan string, closeChan chan struct{}) chan string {
	outChan := make(chan string)

	go func() {
		defer close(outChan)
		for {
			select {
			case <-closeChan:
				return
			default:
			}

			select {
			case value, ok := <-inChan:
				if !ok {
					return
				}
				outChan <- value
			case <-closeChan:
				return
			}
		}
	}()

	return outChan
}

//func main() {
//	inChan := make(chan string)
//	go func() {
//		for i := 0; i < 5; i++ {
//			inChan <- fmt.Sprintf("test: %d", i)
//		}
//	}()
//
//	closeCh := make(chan struct{})
//	go func() {
//		time.Sleep(3 * time.Second)
//		close(closeCh)
//	}()
//
//	for value := range OrDone(inChan, closeCh) {
//		fmt.Println(value)
//	}
//}
