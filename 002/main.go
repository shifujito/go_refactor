package main

import (
	"errors"
	"fmt"
)

func nestJoin(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		concat := s1 + s2
		if s2 == "" {
			return "", errors.New("s1 is empty")
		} else {
			if len(concat) > max {
				return concat[:max], nil
			} else {
				return concat, nil
			}
		}
	}
}

func refactorJoin(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat := s1 + s2
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func main() {
	moji, err := nestJoin("aaa", "bbb", 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(moji)
	moji2, err := refactorJoin("aaa", "bbb", 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(moji2)
}
