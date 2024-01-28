package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	strings := []string{"abc", "saaaaa", "bbbbbb"}

	isLong := func(x string) bool {
		return len(x) >= 4
	}

	filteredInts := Filter(ints, isEven)
	fmt.Println(filteredInts)

	filteredStrings := Filter(strings, isLong)
	fmt.Println(filteredStrings)
}

type fils interface {
	~int | ~string
}

func Filter[T fils](slice []T, cond func(T) bool) (result []T) {
	for _, v := range slice {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}
