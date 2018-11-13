package main

import "fmt"

func main() {
	//简单类型
	var b byte
	var i int
	var ss string
	fmt.Printf(" byte: %v \n int: %v \n string: %v \n", b, i, ss)


	// 复杂类型
	var x []int
	var s struct{}
	var ifa interface{}
	var c chan int
	var m map[string] int

	fmt.Println("[]int:", x)
	fmt.Println("struct:",s)
	fmt.Println("interface: ",ifa)
	fmt.Println("chan int: ",c)
	fmt.Println("map[string]int:",m)

}
