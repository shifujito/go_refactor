package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// No.55 並行処理と並列処理を混同する
	// 並行処理は複数の処理を同時に実行することであり、並列処理は複数の処理を同時に実行することである。

	// No.56 並行処理は常に高速ではない
	// No.56 並行処理は常に高速ではない
	// No.57 チェネルとミューテックスの使い分けに悩む
	// No.58 競合問題を理解していない
	i := 0
	go func() {
		i++
	}()
	go func() {
		i++
	}()
	go func() {
		i++
	}()
	go func() {
		i++
	}()
	go func() {
		i++
	}()
	go func() {
		i++
	}()
	fmt.Println(i) // 0
	// このコードは、iの値をインクリメントするために6つのゴルーチンを起動しているが、iに書き戻す前にfmt.Println(i)が実行されているため、iの値は0のままである。
	// 上記を解決するためにはアトミックにする必要がある。
	var j int64
	go func() {
		atomic.AddInt64(&j, 1)
	}()
	go func() {
		atomic.AddInt64(&j, 1)
	}()

	fmt.Println(j) // 2
	// ミューテックを使用する。
	k := 0
	mutex := sync.Mutex{}

	go func() {
		mutex.Lock()
		k++
		mutex.Unlock()
	}()

	go func() {
		mutex.Lock()
		k++
		mutex.Unlock()
	}()

	fmt.Println(k) // 2

	// Goメモリモデル
}
