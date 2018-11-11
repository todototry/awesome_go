package main

import "fmt"
// panic 之后， recover（）的捕获严重依赖于非 lambda 的函数栈空间
func main() { // call depth 0
	defer func() { // call depth 1  // 含栈空间
		fmt.Println("Now, the panic is still in call depth 01 ", recover())
		func() { // call depth 2
			fmt.Println("Now, the panic is still in call depth 02 ", recover())
			func() { // call depth 3
				fmt.Println("Now, the panic is still in call depth 03 ", recover())
			}()
		}()
	}()

	defer fmt.Println("Now, the panic is in call depth 0 ", recover())

	func() { // call depth 1
		defer fmt.Println("Now, the panic is in call depth 1 ", recover())
		func() { // call depth 2
			//defer func() {
			//	fmt.Println("Now, the panic is in call depth 2 ", recover()) // 含栈空间
			//}()

			fmt.Println("Now, the panic is in call depth 2 ", recover())

			func() { // call depth 3
				defer fmt.Println("Now, the panic is in call depth 3 ", recover())

				panic(1000)
			}()
		}()
	}()
}
