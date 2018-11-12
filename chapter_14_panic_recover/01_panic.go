package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			fmt.Println("6:", recover())
		}()
	}()
	defer func() {
		func() {
			fmt.Println("5:", recover())
		}()
	}()
	func() {
		defer func() {
			fmt.Println("1:", recover())
		}()
	}()
	func() {
		defer fmt.Println("2:", recover())
	}()
	func() {
		fmt.Println("3:", recover())
	}()

	fmt.Println("4:", recover())

	panic(1)

	defer func() {
		fmt.Println("0:", recover()) // never go here
	}()
}
