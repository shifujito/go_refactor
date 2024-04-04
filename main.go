package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
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
	fmt.Println(&dst[0])
	fmt.Println(&src[0])
}

func mistakeAppend() {
	// 間違った追加
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)
	fmt.Println(s1) // [1 2 10]
	fmt.Println(s2) // [2]
	fmt.Println(s3) // [2 10]

	// 正しい追加(s1に追加させない)
	s1 = []int{1, 2, 3}
	sCopy := make([]int, 2)
	copy(sCopy, s1)
	result := append(sCopy, 10)
	fmt.Println(s1)     // [1 2 3]
	fmt.Println(sCopy)  // [1 2]
	fmt.Println(result) // [1 2 10]
}

// スライスとメモリリーク
func sliceAndMemoryLeak() {
	for i := 0; i < 10; i++ {
		// 文字列からbyte配列を作成
		msg := []byte("abcdefghijklmnopqrstuvwxyz")
		fmt.Println(msg)
		badMsg := getBadMessageType(msg)
		fmt.Println(badMsg)
		// newMsgの長さを出力
		fmt.Println(len(badMsg))
		// newMsgの容量を出力
		fmt.Println(cap(badMsg))
		// newMsgの容量を変更
		goodMsg := getGoodMessageType(msg)
		fmt.Println(goodMsg)
		// newMsgの長さを出力
		fmt.Println(len(goodMsg))
		// newMsgの容量を出力
		fmt.Println(cap(goodMsg))
	}
	fmt.Println("-------sliceAndMemoryLeak end-------")
	sliceAndPointer()
}

// msgをスライス化してメッセージ種別を計算する
func getBadMessageType(msg []byte) []byte {
	return msg[:3]
}

func getGoodMessageType(msg []byte) []byte {
	// msg[:3]を返す
	msgType := make([]byte, 3)
	copy(msgType, msg[:3])
	return msgType
}

type Foo struct {
	v []byte
}

func sliceAndPointer() {
	foos := make([]Foo, 1000)
	printAlloc()
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two := keepFirstTwoElementsOnly(foos)
	fmt.Println(two[0].v[0])
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	// foosの最初の2つの要素を保持するスライスを作成
	// return foos[:2]
	// return foos[:2:2]
	res := make([]Foo, 2)
	copy(res, foos)
	return res
}
func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
}

// 非効率なマップの初期化
func initMap() {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	fmt.Println(m)
}

// マップとメモリリーク
func mapAndMemoryLeak() {
	n := 1000000
	m := make(map[int][128]byte)
	printAlloc()
	for i := 0; i < n; i++ {
		m[i] = [128]byte{}
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

// 値の比較誤り
type costumer struct {
	id         string
	operations []float64
}

func compareValue() {
	cst := costumer{id: "1", operations: []float64{1.0}}
	cst2 := costumer{id: "1", operations: []float64{1.0}}
	fmt.Println(reflect.DeepEqual(cst, cst2))
}

// rangeループで要素がコピーされることを無視する
type account struct {
	balance float64
}

func rangeLoop() {
	// 更新されない
	accounts := []account{{100.0}, {200.0}, {300.0}}
	for _, acc := range accounts {
		acc.balance += 1000
	}
	fmt.Println(accounts)
	// 更新されるパターン1
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)
	// 更新されるパターン2
	accounts2 := []*account{{100.0}, {200.0}, {300.0}}
	for _, acc := range accounts2 {
		acc.balance += 1000
	}
	fmt.Println(accounts2)
}

// rangeループでの引数の評価方法を無視する
func ignoreRangeLoop() {
	s := []int{0, 1, 2}
	for range s {
		s = append(s, 3)
	}
	fmt.Println(s)

	// 無限ループになる
	// sb := []int{0, 1, 2}
	// for i := 0; i < len(sb); i++ {
	// 	sb = append(sb, 3)
	// }

	ch1 := make(chan int, 3)
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	ch2 := make(chan int, 3)
	go func() {
		ch2 <- 10
		ch2 <- 11
		ch2 <- 12
	}()

	ch := ch1
	for v := range ch {
		// ch1に対して評価される。
		fmt.Println(v)
		ch = ch2
	}

	a := [3]int{1, 2, 3}
	for i, v := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println("v1", v)
			fmt.Println("v1", a[2])
		}
	}
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	for i, v := range &b {
		b[2] = 100
		if i == 2 {
			fmt.Println("v2", v)
		}
	}
	fmt.Println(b)
}

type Customer struct {
	ID      string
	Balance float64
}
type Store struct {
	m map[string]*Customer
}

func (s *Store) storeCustomers(customers []Customer) {
	for _, c := range customers {
		c := c
		fmt.Printf("%p\n", &c)
		s.m[c.ID] = &c
	}
}

func ignoreRangeLoopPointer() {
	s := Store{m: make(map[string]*Customer)}
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 100},
		{ID: "2", Balance: 200},
		{ID: "3", Balance: 300},
	})
	for k, v := range s.m {
		fmt.Println(k, v)
	}
}

func mapOrder() {
	// 誤った例
	m := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	fmt.Println(m)
	// 正しい例
	m2 := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	m3 := make(map[int]bool, len(m2))
	for k, v := range m2 {
		m3[k] = v
		if v {
			m3[10+k] = true
		}
	}
	fmt.Println(m3)
}

// No.34 break文の仕組みを無視する
func ignoreBreak() {
	// ループ内でswitchやselectと組み合わせてbreakを使う
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i)
		switch i {
		default:
		case 2:
			break
		}
	}
	// break文は一番内側のfor文、switch文、select文を抜ける
	// switch文ではなく、for文を抜ける
loop:
	for i := 5; i < 10; i++ {
		fmt.Printf("%d\n", i)
		switch i {
		default:
		case 7:
			break loop
		}
	}
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
		"no26": sliceAndMemoryLeak,
		"no27": initMap,
		"no28": mapAndMemoryLeak,
		"no29": compareValue,
		"no30": rangeLoop,
		"no31": ignoreRangeLoop,
		"no32": ignoreRangeLoopPointer,
		"no33": mapOrder,
		"no34": ignoreBreak,
	}
	f, exists := funcs[name]
	if !exists {
		return nil, fmt.Errorf("関数が見つかりません: %s", name)
	}
	return f, nil
}
