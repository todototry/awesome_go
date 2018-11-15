package main

import "fmt"

// TODO: 本函数返回一个函数： 该函数返回一个返回值是.....
func fun() func() func() func() func() func() int {
	return func() func() func() func() func() int {
		return func() func() func() func() int{
			return func() func() func() int{
				return func() func() int{
					return  func() int{
						return 1
					}
				}
			}
		}
	}

}


func main() {
	fmt.Printf("%T",fun())
}
