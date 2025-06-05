package main

import (
	"fmt"
)

func PrintSlice[T any](s []T) {

	for _, v := range s {
		fmt.Println(v)
	}

}
func main() {
	a := []int{1, 2, 3, 4, 5}
	PrintSlice(a)

}
