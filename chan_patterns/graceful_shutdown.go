package chan_patterns

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func shutdown() {
	interruptChan := make(chan os.Signal)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer func() {
			ticker.Stop()
			wg.Done()
		}()

		for {
			select {
			case <-interruptChan:
				log.Println("shutdown")
				return
			default:
			}

			select {
			case <-interruptChan:
				log.Println("shutdown")
				return
			case <-ticker.C:
				log.Println("do work")
			}
		}
	}()

	wg.Wait()
	log.Println("App was stopped")
}
