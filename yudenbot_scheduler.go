//yudengo_scheduler
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Executor struct {
	Name    string
	Fnc     func(context.Context) error
	Tick    time.Duration
	_ticker *time.Ticker
}

func Schedule(ctx context.Context, execList []Executor) (err error) {
	log.Println("scheduler start")
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGINT,
	)

	ctx, cancelFnc := context.WithCancel(ctx)
	chStop := make(chan int, 1)
	go func(stopTimer chan int, ctx context.Context) {
		for n := range execList {
			execList[n]._ticker = time.NewTicker(execList[n].Tick)
			defer execList[n]._ticker.Stop()
		}

		masterTick := time.NewTicker(1 * time.Second)
		defer masterTick.Stop()

	LOOP:
		for {
			for _, e := range execList {
				select {
				case <-e._ticker.C:
					log.Println("tick : ", e.Name)
					go e.Fnc(ctx)
				case <-stopTimer:
					log.Println("Timer stop.")
					break LOOP
				case <-masterTick.C:
					continue
				}
			}
		}
		log.Println("timerfunc end.")
	}(chStop, ctx)

	sigCh := <-signalCh
	cancelFnc()
	// catch os signal
	log.Println("!! catch signal !! : ", sigCh)

	chStop <- 0 // stop ticker
	close(chStop)
	return nil
}
