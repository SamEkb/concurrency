package sync_primitives

import "sync"

type Semaphore struct {
	size  int
	count int
	sync  *sync.Cond
}

func NewSemaphore(size int) *Semaphore {
	m := &sync.Mutex{}
	return &Semaphore{
		size: size,
		sync: sync.NewCond(m),
	}
}

func (s *Semaphore) Acquire() {
	s.sync.L.Lock()
	defer s.sync.L.Unlock()

	if s.count >= s.size {
		s.sync.Wait()
	}

	s.count++
}

func (s *Semaphore) Release() {
	s.sync.L.Lock()
	defer s.sync.L.Unlock()

	s.count--
	s.sync.Signal()
}
