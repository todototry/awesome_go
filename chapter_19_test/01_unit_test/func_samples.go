package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func addSlice(a ...int) int {
	n := 0
	for _, v := range a {
		n += v
	}

	return n
}

func adjoin(a, b string) string {
	return a + b
}

func main() {
	//fmt.Println(add)

	fmt.Println(add(1, 3))
	fmt.Println(addSlice(1, 3, 5))
}
