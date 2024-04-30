package main

import (
	"fmt"
)

func main() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	close(ch)

	k := 0
	chk := make(chan struct{}, 1)
	go func() {
		k = 1
		<-chk
	}()
	chk <- struct{}{}
	fmt.Println("k", k)

	j := 0
	chj := make(chan struct{})
	go func() {
		j = 1
		<-chj
	}()
	chj <- struct{}{}
	fmt.Println("j", j)
}
