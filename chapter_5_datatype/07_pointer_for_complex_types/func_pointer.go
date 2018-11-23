package main

import "fmt"

// 参考: https://stackoverflow.com/questions/3601796/can-we-have-function-pointers-in-google-go

type ArithOp func(int, int) int

func main() {
	calculate(Plus)
	calculate(Minus)
	calculate(Multiply)
}

func calculate(fp func(int, int) int) {
	ans := fp(3, 2)
	fmt.Printf("\n%v\n", ans)
}

// This is the same function but uses the type/fp defined above
//
// func calculate (fp ArithOp) {
//     ans := fp(3,2)
//     fmt.Printf("\n%v\n", ans)
// }

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}
