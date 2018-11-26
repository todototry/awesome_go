package main

import (
	"fmt"
	"time"
)

func normal_select(){
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

func no_buffer_select(){
	// 注意添加 chan 的 size， 否则是 xchan 就是 nil。
	var xchan = make(chan int, 1)

		go func() {
			num := 1
			for {
				fmt.Println(num)
				time.Sleep(1e9)

				num++

				if num == 3 {
					break
				}
			}
			close(xchan)
		}()

		select {
			case <- xchan:  //本行一定要添加，否则出现 all go routines are sleep, deadlock .
		}
}

/*
哪些情形下初始化的 value 是 nil?

// This following line doesn't compile.
	v := nil

// There must be sufficient information for compiler
// to deduce the type of a nil value.
_ = (*struct{})(nil)
_ = []int(nil)
_ = map[int]bool(nil)
_ = chan string(nil)
_ = (func())(nil)
_ = interface{}(nil)

// This lines are equivalent to the above lines.
var _ *struct{} = nil
var _ []int = nil
var _ map[int]bool = nil
var _ chan string = nil
var _ func() = nil
var _ interface{} = nil

//

*/

func main() {

	normal_select()
	no_buffer_select()


	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)


}
