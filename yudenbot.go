/*
YudenBot is supporter of infra-workshop(インフラ勉強会).

What is infra-workshop(インフラ勉強会) ?

Infra-workshop is japanese online community for study computer infrastructure.
(infra-workshop writes as "インフラ勉強会" in Japanese.)

More information

https://wp.infra-workshop.tech/ (Japanese/日本語)
*/
package main

// とりあえずローカルで動くように

import (
	"context"
	"log"
)

func Yudenbot(ctx context.Context, execList []Executor) {
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
	Schedule(ctx, execList)
	log.Println("Yuden-Bot End.")
}

// executer-d
// discordにpost

// executer-t
// twitterにpost
