package chan_patterns

import (
	"sync"
)

type Barrier struct {
	maxSize  int
	count    int
	mx       sync.Mutex
	beforeCh chan int
	afterCh  chan int
}

func NewBarrier(maxSize int) *Barrier {
	return &Barrier{
		maxSize:  maxSize,
		beforeCh: make(chan int, maxSize),
		afterCh:  make(chan int, maxSize),
	}
}

func (b *Barrier) Before() {
	b.mx.Lock()

	b.count++
	// Когда счетчик будет равен количеству горутин, зайдем, запишем в канал и отпустим заблокированные горутины
	if b.count == b.maxSize {
		for i := 0; i < b.maxSize; i++ {
			b.beforeCh <- 1
		}
	}

	b.mx.Unlock()
	// Горутины блокируются и ждут, когда ктонибудь запишет в канал, чтобы прочитать
	<-b.beforeCh
}

func (b *Barrier) After() {
	b.mx.Lock()

	// Здесь наоборот декремент, когда он выполнен и счетчик 0, то пишем в канал и отпускем горутины
	b.count--
	if b.count == 0 {
		for i := 0; i < 3; i++ {
			b.afterCh <- 1
		}
	}

	b.mx.Unlock()
	// Блокируем горутины и ждем, когда ктонибудь запишет в канал, чтобы прочитать
	<-b.afterCh
}

//func main() {
//	wg := sync.WaitGroup{}
//
//	job := func() {
//		fmt.Println("job")
//	}
//
//	work := func() {
//		fmt.Println("work")
//	}
//
//	barrier := NewBarrier(3)
//	for i := 0; i < 3; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for i := 0; i < 3; i++ {
//				//Блокируем горутины
//				barrier.Before()
//				//Выполняем, то количество горутин, которое проитерировалось
//				job()
//				//Блокируем горутины
//				barrier.After()
//				//Выполняем, то количество горутин, которое проитерировалось
//				work()
//			}
//		}()
//	}
//
//	wg.Wait()
//}
