package chan_patterns

type Future struct {
	result chan interface{}
}

func NewFuture(task func() interface{}) *Future {
	if task() == nil {
		return nil
	}

	future := &Future{result: make(chan interface{})}

	go func() {
		defer close(future.result)
		future.result <- task()
	}()

	return future
}

func (f *Future) Get() interface{} {
	return <-f.result
}

//func main() {
//	callback := func() interface{} {
//		return "result"
//	}
//
//	future := NewFuture(callback)
//	result := future.Get()
//	fmt.Println(result)
//}
