package main

import "fmt"

func main() {
	//简单类型
	var b byte
	var i int
	var ss string
	fmt.Printf(" byte: %v \n int: %v \n string: %v \n", b, i, ss)

	// 复杂类型
	/*
		Types which zero values can be represented with nil
		The zero values of the following types can be represented with nil.

			Type (T)	Size Of T(nil)
			pointer	1 word
			slice	3 words
			map	1 word
			channel	1 word
			function	1 word
			interface	2 words
	*/
	var x []int
	var s struct{}
	var ifa interface{}
	var c chan int
	var m map[string]int

	fmt.Println("[]int:", x)
	fmt.Println("struct:", s)
	fmt.Println("interface: ", ifa)
	fmt.Println("chan int: ", c)
	fmt.Println("map[string]int:", m)

}
