package main

import "fmt"

func defer_in_block() {
	i := 0
	{
		defer fmt.Println("defer in code block begin")
		fmt.Println("in block")
		defer fmt.Println("defer in code block end")
	}

	defer 	fmt.Println("defer fmt in function start 1")
	fmt.Println("fmt in function 1")
	defer 	fmt.Println("defer fmt in function end  1")


	if i == 0 {
		defer 	fmt.Println("defer fmt in if{} block start")
		fmt.Println("fmt in if()")
		defer 	fmt.Println("defer fmt in if {} block end")
	}

	defer 	fmt.Println("defer fmt in function start 2")
	fmt.Println("fmt in function  2")
	defer 	fmt.Println("defer fmt in function end 2")

	}

func main() {
	defer_in_block()
}
