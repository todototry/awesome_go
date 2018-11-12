package main

import "fmt"

func demo(recoverHighestPanicAtFirst bool) {
	fmt.Println("====================")
	defer func() {
		if !recoverHighestPanicAtFirst{
			// recover panic 1
			defer fmt.Println("panic", recover(), "is recovered")
		}
		defer func() {
			// recover panic 2
			fmt.Println("panic", recover(), "is recovered")
		}()
		if recoverHighestPanicAtFirst {
			// recover panic 1
			defer fmt.Println("panic", recover(), "is recovered")
		}
		defer fmt.Println("now, two active panics coexist")
		panic(2)
	}()
	panic(1)
}

func main() {
	demo(true)
	demo(false)
}
