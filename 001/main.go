package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func BadScope() {
	var badScope string
	flag := true
	if flag {
		badScope, err := "aaa", xerrors.New("")
		if err != nil {
			fmt.Println("err") //err
		}
		fmt.Println(badScope) // aaa
	} else {
		badScope, err := "bbb", xerrors.New("")
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println(badScope)
	}
	fmt.Println(badScope) // null
}

func GoodScope() {
	var goodScope string
	flag := true
	if flag {
		scope, err := "aaa", xerrors.New("")
		if err != nil {
			fmt.Println("err") //err
		}
		fmt.Println(scope) // aaa
		goodScope = scope
	} else {
		scope, err := "bbb", xerrors.New("")
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println(scope)
		goodScope = scope
	}
	fmt.Println(goodScope) // aaa
}

func main() {
	BadScope()
	GoodScope()
}
