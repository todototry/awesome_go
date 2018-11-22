package main

import (
	"fmt"
)

func main() {
	type NamedParamsInit struct {
		_    struct{}
		name string
		age  int
	}

	x := NamedParamsInit{name: "fan", age: 10}

	// failed
	// y := NamedParamsInit{"", "dong", 1}

	y := NamedParamsInit{struct{}{}, "dong", 1}
	fmt.Println(x, y)

}
