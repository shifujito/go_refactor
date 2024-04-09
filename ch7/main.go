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

func getFunc(i *int) (func(), error) {
	funcs := map[string]func(){
		"1":  func() { fmt.Println("Function 1") },
		"48": panicFunc,
		"49": wrapFoo,
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
