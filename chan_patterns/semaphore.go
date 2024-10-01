package chan_patterns

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, size),
	}
}

func (s *Semaphore) Acquire() {
	if s.ch == nil {
		return
	}
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	if s.ch == nil {
		return
	}
	<-s.ch
}
