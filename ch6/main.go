package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// No.42 どちらのレシーバ型を使うべきか
// 値レシーバーでは、値のコピーを作成しメソッドに渡す
// ポインタレシーバーでは、ポインタを渡すため、値をコピーする必要がない

// レシーバがポインタでなければいけない場合
// 1. メソッドがレシーバの値を変更する必要がある場合
// 2. レシーバがコピーできないフィルドを持っている場合(syncなど)

// レシーバがポインタであるべき場合
// 1. レシーバが大きい場合

// レシーバが値であるべき場合
// 1. メソッドがレシーバの値を変更しない場合
// 2. レシーバの不変性を保証する場合　マップ、関数、チャネルなど
// 3. 変更する必要のないスライスの場合

type customType struct {
	balance float64
	data    *data
}

type data struct {
	balance float64
}

func (c *customType) addBalance(amount float64) {
	c.balance += amount
}

func (c customType) add2Balance(amount float64) {
	c.balance += amount
}

func (c customType) addDataBalance(amount float64) {
	c.data.balance += amount
}

func whichReceiverType() {
	c := customType{balance: 100, data: &data{balance: 100}}
	c.addBalance(100)
	fmt.Println(c.balance) // 200

	// 値レシーバーを使っているので、更新されない。
	c.add2Balance(100)
	fmt.Println(c.balance) // 300ではなく、200

	c.addDataBalance(100)
	fmt.Println(c.data.balance) // 200
}

// No.43 名前付き結果パラメータを使わない
// 名前付き結果パラメータは、関数の戻り値に名前をつけることができますが、Goではあまり使われていない。
// 以下の理由から、名前付き結果パラメータを使わない方が良い。
// 1. 名前付き結果パラメータは、関数の可読性を損なう
// 2. 名前付き結果パラメータは、関数の戻り値が変更される可能性があるため、関数の可読性を損なう
// 3. 名前付き結果パラメータは、関数の戻り値が複数ある場合、名前付き結果パラメータを使うと、戻り値の順序が変わる可能性があるため、可読性を損なう

func useNamedResult() {
	c := useNamedResult2(2)
	fmt.Println(c)
}

func useNamedResult2(a int) (b int) {
	// 引数なしでreturnすると、名前付き結果パラメータが返される
	b = a
	return
}

// 名前付き結果パラメータが推奨される場合
// locatorインターフェスは非公開なので、ドキュメントが不要。コードを読むだけで推測できる。
// type locator interface {
// 	getCoordinates(address string) (float32, float32, error)
// }

// type locator2 interface {
// 	getCoordinates(address string) (lat, lng float32, err error)
// }

// func (l loc) getCoordinates(address string) (lat, lng float32, err error) {
// 	// do something
// }

// No.44 名前付きパラメータによる意図しない副作用

func namedResult() {
	fmt.Println("--- namedResult ---")
}

// No.45 nilレシーバを返す
type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

func validate() {
	var m *MultiError
	name := "1"
	if name == "" {
		m = &MultiError{}
		m.Add(errors.New("error1"))
	} else if name == "test" {
		m = &MultiError{}
		m.Add(errors.New("error2"))
	}
	fmt.Println(m)
	var foo *Foo
	fmt.Println(foo.Bar())
	fmt.Println(foo)

}

type Foo struct{}

func (f *Foo) Bar() string {
	return "Bar"
}

// No.46 ファイル名を関数の入力として使う
func useFileNameAsInput() {
	fileName := "test.txt"
	count, err := countEmptyLinesInFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Empty lines: %d\n", count)
}

func countEmptyLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			// do something
		}
	}
	return 0, nil
}

// No.47 deferの引数やレシーバの評価方法を無視している
const (
	StatusSuccess = "success"
	StatusFailed  = "failed"
	StatusFoo     = "foo"
)

func deferEvaluation() {
	var status string
	// Goはdeferを呼び出すときに引数を評価する
	// これの解決策はポインタを渡すこと
	defer fmt.Println(status)
}

func getFunc(i *int) (func(), error) {
	funcs := map[string]func(){
		"1":  func() { fmt.Println("Function 1") },
		"42": whichReceiverType,
		"43": useNamedResult,
		"44": namedResult,
		"45": validate,
		"46": useFileNameAsInput,
		"47": deferEvaluation,
	}
	f, ok := funcs[fmt.Sprintf("%d", *i)]
	if !ok {
		return nil, fmt.Errorf("Function not found")
	}
	return f, nil
}

func main() {
	i := flag.Int("i", 0, "integer flag")
	flag.Parse()
	fmt.Printf("No of arguments: %d\n", *i)

	f, err := getFunc(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	f()
}
