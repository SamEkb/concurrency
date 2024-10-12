package chan_patterns

import (
	"fmt"
	"time"
)

func DoWork(inputCh <-chan struct{}) <-chan struct{} {
	outputCh := make(chan struct{})

	go func() {
		ticker := time.NewTicker(time.Millisecond * 100)
		defer func() {
			close(outputCh)
			ticker.Stop()
		}()

		for {
			select {
			case <-inputCh:
				fmt.Println("Channel is closed, that's all")
				return
			case <-ticker.C:
				fmt.Println("Do some work")
			}
		}
	}()
	return outputCh
}

//func main() {
//	closeCh := make(chan struct{})
//	closeDoneCh := DoWork(closeCh)
//	time.Sleep(time.Second)
//	close(closeCh)
//	<-closeDoneCh
//}
