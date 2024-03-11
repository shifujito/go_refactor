package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

var funcs = map[string]func(){
	"no17": no17,
	"no18": no18,
	"no19": no19,
}

func main() {
	i := flag.Int("i", 0, "int flag")
	flag.Parse()
	fmt.Printf("param -i : %d\n", *i)

	funcNo := "no" + strconv.Itoa(*i)
	if f, exists := funcs[funcNo]; exists {
		fmt.Printf("run method : no%d\n", *i)
		f() // マップから取得した関数を実行します。
	} else {
		fmt.Println("関数が見つかりません:", funcNo)
	}
}

func no17() {
	sum := 100 + 010
	fmt.Println(sum)
	os.OpenFile("memo.md", os.O_CREATE, 0644)
}

func no18() {
	var counter int32 = math.MaxInt32
	counter++
	fmt.Println(counter) //-2147483648
}

func no19() {
	var n float32 = 1.0001
	fmt.Println(n)
	fmt.Println(n * n)
}
