package main

import (
	"errors"
	"flag"
	"fmt"
)

// No.48 パニックを発生させる
// Goでは、プログラマーエラーはpanicを使って表現することが推奨されている。
func panicFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic("panic!")
}

// No.49 エラーをラップすべきときを無視する。
// エラーラッピングはエラーエラーラップして元のエラーも利用するようにする。
func foo() error {
	err := errors.New("original error")
	return fmt.Errorf("an error occurred: %w", err)
}

func wrapFoo() {
	err := foo()
	fmt.Println(err) // "an error occurred: original error"

	originalErr := errors.Unwrap(err)
	fmt.Println(originalErr) // "original error"
}

// No.50 エラー型を不正確に検査する
type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount(ID string) (float32, error) {
	if len(ID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", ID)
	}

	return 100.0, nil
}

func nm50() {
	amount, err := getTransactionAmount("12345")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(amount)
}

// No.51 エラー値を不正確に検査する
func nm51() {
	var ErrFoo = errors.New("foo error")
	fmt.Println(ErrFoo)
}

type Route struct {
	Lat, Lng float64
}

// No.52 エラーを二度処理する
func nm52() {
	route, err := validate(91.0, 181.0)
	fmt.Println(route, err)
}

func validate(lat, lng float64) (Route, error) {
	if lat > 90.0 || lng < -90.0 {
		return Route{}, fmt.Errorf("invalid latitude or longitude")
	}
	return Route{Lat: lat, Lng: lng}, nil
}

func getFunc(i *int) (func(), error) {
	funcs := map[string]func(){
		"1":  func() { fmt.Println("Function 1") },
		"48": panicFunc,
		"49": wrapFoo,
		"50": nm50,
		"51": nm51,
		"52": nm52,
	}
	f, ok := funcs[fmt.Sprintf("%d", *i)]
	if !ok {
		return nil, fmt.Errorf("function not found")
	}
	return f, nil
}

func main() {
	i := flag.Int("i", 0, "int flag")
	flag.Parse()

	f, err := getFunc(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	f()
}
