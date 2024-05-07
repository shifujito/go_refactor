package main

import (
	"encoding/json"
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

func no76() {
	// No.76 time.Afterとメモリリーク
	// time.Afterは、チェンルを返し、指定された時間が経過後にチャネルに値を送信する。
	// time.Afterをループで使用するとメモリリークが発生する。
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("tick")
		}
	}
}

// Json処理の間違い
type Event struct {
	ID        int
	time.Time // 埋め込みフィールド
}

func no77() {

	event := Event{
		ID:   1,
		Time: time.Now(),
	}
	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(b))
	// time.Timeは構造体であり、MarshalJSONメソッドを持っている。
	// そのため、time.Timeを埋め込んだ構造体をJSONに変換すると、time.Timeのフィールドが直接JSONに変換される。
	// 解決策1. json.Marshalerを実装する。
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int
		Time time.Time
	}{
		ID:   e.ID,
		Time: e.Time,
	})
}

func main() {
	no77()
}
