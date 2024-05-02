package main

import (
	"fmt"
	"time"
)

func no75() {
	// No.75 誤った時間の長さを提供する。
	// 標準ライブラリには、time.Durationを受け取る一般的な関数やメソッドが用意されているが、time.Durationはint64型のエイリアスである。
	ticker := time.NewTicker(1 * time.Second)
	stop := time.After(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-stop:
			ticker.Stop()
			fmt.Println("finish")
			return
		}
	}
}

func main() {
	no75()
}
