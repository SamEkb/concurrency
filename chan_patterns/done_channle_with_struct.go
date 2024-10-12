package chan_patterns

import (
	"fmt"
	"time"
)

type Worker struct {
	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewWorker() *Worker {
	worker := &Worker{
		closeCh:     make(chan struct{}),
		closeDoneCh: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer func() {
			close(worker.closeDoneCh)
			ticker.Stop()
		}()

		for {
			select {
			case <-worker.closeCh:
				fmt.Println("Channel was closed")
				return
			default:
			}

			select {
			case <-worker.closeCh:
				fmt.Println("Channel was closed")
				return
			case <-ticker.C:
				fmt.Println("Do some work")
			}
		}
	}()

	return worker
}

func (w *Worker) shutdown() {
	close(w.closeCh)
	<-w.closeDoneCh
}

//func main() {
//	worker := NewWorker()
//	fmt.Println("Work was started")
//	time.Sleep(time.Second * 5)
//	worker.shutdown()
//}
