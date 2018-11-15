package main

import "fmt"

func scope(){

	var x int = 1

	{
		var y int = 2
		fmt.Println(x)
		fmt.Println(y)
	}
	fmt.Println(x)
	// TODO: read y out of scope is invalid.
	// fmt.Println(x, y)
}


func main() {
	scope()
}
