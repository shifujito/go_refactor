package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	i := flag.Int("i", 0, "int flag")
	flag.Parse()
	fmt.Printf("param -i : %d\n", *i)

	funcNo := "no" + strconv.Itoa(*i)
	f, err := getFunc(funcNo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("run method : no%d\n", *i)
	f() // マップから取得した関数を実行します。
}

func addNumbers() {
	sum := 100 + 010
	fmt.Println(sum)
	os.OpenFile("memo.md", os.O_CREATE, 0644)
}

func incrementCounter() {
	var counter int32 = math.MaxInt32
	counter++
	fmt.Println(counter) //-2147483648
}

func printFloat() {
	var n float32 = 1.0001
	fmt.Println(n)
	fmt.Println(n * n)
}

func initSlice() {
	sts := []string{"a", "b", "c"}
	// 非効率なスライスの初期化
	bars := make([]string, 0)
	for _, foo := range sts {
		bars = append(bars, foo)
	}
	fmt.Println(bars)

	// 効率的なスライスの初期化
	n := len(sts)
	bars = make([]string, 0, n)
	for _, foo := range sts {
		bars = append(bars, foo)
	}
	fmt.Println(bars)

	// さらに効率化する
	bars = make([]string, n)
	for i, foo := range sts {
		bars[i] = foo
	}
	fmt.Println(bars)
}

func nilSliceAndEmptySlice() {
	// nilスライス
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)

	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}

func handleOperations() {
	id := ""
	operations := getOperation(id)
	if operations == nil {
		fmt.Println("操作がありません")
		return
	}
	fmt.Println(operations)
}

func getOperation(id string) []float32 {
	operations := make([]float32, 0)
	if id == "" {
		// return operationsではなく、return nilを返す
		return nil
		// return operations
	}
	// 要素を追加
	operations = append(operations, 2.0)
	return operations
}

func mistakeCopy() {
	// 間違ったコピー
	src := []int{1, 2, 3}
	var dst []int
	copy(dst, src)
	fmt.Println(dst)

	// 正しいコピー
	src = []int{1, 2, 3}
	dst = make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst)
}

func mistakeAppend() {
	// 間違った追加
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)
	fmt.Println(s3)

	// 正しい追加
	// dst = []int{1, 2, 3}
	// dst = append(dst, 4)
	// fmt.Println(dst)
}

func getFunc(name string) (func(), error) {
	funcs := map[string]func(){
		"no17": addNumbers,
		"no18": incrementCounter,
		"no19": printFloat,
		"no21": initSlice,
		"no22": nilSliceAndEmptySlice,
		"no23": handleOperations,
		"no24": mistakeCopy,
		"no25": mistakeAppend,
	}

	f, exists := funcs[name]
	if !exists {
		return nil, fmt.Errorf("関数が見つかりません: %s", name)
	}

	return f, nil
}
