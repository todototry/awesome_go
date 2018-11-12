package main

import "fmt"

func main() { // call depth 0

	defer func() { // call depth 1, this can absorb panic 3 from depth 0.
		fmt.Println("with out this, program will crash, pick the panc 3: see: ", recover())
	}()

	defer fmt.Println("program will crash, for panic 3 is stll active")

	defer func() { // call depth 1
		defer func() { // call depth 2
			fmt.Println( recover() ) // 6
		}()

		// The depth of panic 3 is 0, and the depth of panic 6 is 1.
		defer fmt.Println("now, there are two active panics: 3 and 6, but since we are running at depth 1 , so we can only get panic  3 which occur at depth 0. ")
		defer panic(6) // will suppress panic 5
		defer panic(5) // will suppress panic 4
		panic(4) // will not suppress panic 3,
		// for they have differrent depths.
		// The depth of panic 3 is 0.
		// The depth of panic 4 is 1.
	}()

	defer fmt.Println("now, only panic 3 is active")
	defer panic(3) // will suppress panic 2
	defer panic(2) // will suppress panic 1
	panic(1)
}
