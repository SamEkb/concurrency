package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[1]
	}

	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-doneCh:
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:])...):
			}
		}
	}()

	return doneCh
}

func main() {
	after := func(duration time.Duration) <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			time.Sleep(duration)
			close(ch)
		}()
		return ch
	}

	startTime := time.Now()

	<-or(
		after(time.Second),
		after(time.Millisecond*100),
		after(time.Millisecond*1000),
		after(time.Minute),
		after(time.Second*2),
		after(time.Millisecond*200),
		after(time.Millisecond*500),
	)
	fmt.Printf("Total time: %s", time.Since(startTime))
}
