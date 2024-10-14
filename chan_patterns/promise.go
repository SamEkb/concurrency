package chan_patterns

type Promise struct {
	waitChan chan struct{}
	value    interface{}
	err      error
}

func NewPromise(task func() (interface{}, error)) *Promise {
	if task == nil {
		return nil
	}

	promise := &Promise{
		waitChan: make(chan struct{}),
	}

	go func() {
		defer close(promise.waitChan)
		promise.value, promise.err = task()
	}()

	return promise
}

func (p *Promise) Then(success func(interface{}), error func(error)) {
	<-p.waitChan

	if p.err == nil {
		success(p.value)
	} else {
		error(p.err)
	}
}

//func main() {
//	callback := func() (interface{}, error) {
//		return "error", errors.New("error")
//	}
//
//	promise := NewPromise(callback)
//	promise.Then(
//		func(value interface{}) {
//			fmt.Println("success", value)
//		},
//		func(err error) {
//			fmt.Println("error", err.Error())
//		})
//}
