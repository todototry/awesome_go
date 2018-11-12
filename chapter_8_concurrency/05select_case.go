package main

import "fmt"

func main() {
	var c1 = make(chan int, 1)
	var c2 = make(chan int, 1)
	var c3 = make(chan int, 1)

	var i1, i2 int

	c1 <- 1

	c3 <- 3

	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3\n")
		} else {
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Println("no communication\n")
	}
}
