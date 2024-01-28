package redis

import "fmt"

func init() {
	fmt.Print("init 1")
}

func Store(key, value string) error {
	return nil
}
