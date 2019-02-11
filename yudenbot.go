/*
YudenBot is supporter of infra-workshop(インフラ勉強会).

What is infra-workshop(インフラ勉強会) ?

Infra-workshop is japanese online community for study computer infrastructure.
(infra-workshop writes as "インフラ勉強会" in Japanese.)

More information

https://wp.infra-workshop.tech/ (Japanese/日本語)
*/
package yudenbot

// とりあえずローカルで動くように

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Executor struct {
	name    string
	fnc     func(context.Context) error
	tick    time.Duration
	_ticker *time.Ticker
}

func Yudenbot(execList []Executor) {
	log.Print("run Yuden-Bot")

	// updater
	// Wordpressから情報Get
	// 書き出す
	// 30分ごと程度

	// fetcher
	// 読み出し
	// 時刻チェック
	// execute()
	// 1分毎
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGINT,
	)

	chStop := make(chan int, 1)
	go func(stopTimer chan int) {
		for n := range execList {
			execList[n]._ticker = time.NewTicker(execList[n].tick)
			defer execList[n]._ticker.Stop()
		}

		masterTick := time.NewTicker(1 * time.Second)
		defer masterTick.Stop()

		var ctx context.Context
	LOOP:
		for {
			for _, e := range execList {
				select {
				case <-e._ticker.C:
					log.Println("tick : ", e.name)
					go e.fnc(ctx)
				case <-stopTimer:
					log.Println("Timer stop.")
					break LOOP
				case <-masterTick.C:
					continue
				}
			}
		}
		log.Println("timerfunc end.")
	}(chStop)

	sigCh := <-signalCh
	// catch os signal
	log.Println("!! catch signal !! : ", sigCh)

	chStop <- 0 // stop ticker
	close(chStop)
	log.Println("Yuden-Bot End.")
}

// executer-d
// discordにpost

// executer-t
// twitterにpost
