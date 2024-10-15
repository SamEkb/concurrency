package chan_patterns

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type ErrGroup struct {
	err    unsafe.Pointer
	wg     sync.WaitGroup
	doneCh chan struct{}
}

func NewErrGroup() *ErrGroup {
	return &ErrGroup{
		doneCh: make(chan struct{}),
	}
}

func (e *ErrGroup) Go(task func() error) {
	select {
	case _, ok := <-e.doneCh:
		if !ok {
			return
		}
	default:
	}

	e.wg.Add(1)
	go func() {
		defer e.wg.Done()

		select {
		case <-e.doneCh:
			return
		default:
			if err := task(); err != nil {
				newPtr := unsafe.Pointer(&err)
				if atomic.CompareAndSwapPointer(&e.err, nil, newPtr) {
					close(e.doneCh)
				}
			}
		}
	}()
}

func (e *ErrGroup) Wait() error {
	e.wg.Wait()
	if err := atomic.LoadPointer(&e.err); err != nil {
		return *(*error)(err)
	} else {
		return nil
	}
}

//func main() {
//	errGroup := NewErrGroup()
//	errGroup.Go(func() error {
//		fmt.Println("Do some job")
//		return nil
//	})
//
//	errGroup.Go(func() error {
//		fmt.Println("Do wrong job")
//		return errors.New("error")
//	})
//
//	time.Sleep(time.Second)
//
//	for i := 0; i < 5; i++ {
//		errGroup.Go(func() error {
//			fmt.Println("Do some job after error")
//			return nil
//		})
//	}
//
//	if err := errGroup.Wait(); err != nil {
//		fmt.Println(err.Error())
//	}
//}
