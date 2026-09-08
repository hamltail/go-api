package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// goを付けると、taskをgoroutineとして開始する。
	// mainはtaskの終了を待たず、すぐ次の行へ進む。
	go task("A")

	// Aが動いている間も、mainは別の処理を続けられる。
	for i := 1; i <= 3; i++ {
		fmt.Println("main", i)
		time.Sleep(500 * time.Millisecond)
	}

	// mainが終了するとプログラム全体が終了し、
	// 実行中のgoroutineも途中で終了してしまう。
	// 今回は動作確認のため、mainを終了させないよう少し待つ。
	// ※ Sleepで待つのはデモ用。本来の同期方法ではない。
	time.Sleep(2 * time.Second)
}
