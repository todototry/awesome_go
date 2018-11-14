package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	x := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			<-c
		}
	}()



	// go routine 无需
	for _,z := range []int{1,2,3,4,5,6,7,8}{
		go func( int) {
			select {
			case i := <-x:
				fmt.Println("out:", i)
			case x <- z:
				fmt.Println("in:", z)
			}

		}(z)

	}

	// 由于 go routine 的执行启动时间不能保证，通常会在 for 执行完成后启动，因此，如果不传入具体数字 z, 那么会使用外部的 z 变量, 也就是最后一个 z 值.
	for _,z := range []int{1,2,3,4,5,6,7,8}{
		go func(z int) {
			select {
			case i := <-x:
				fmt.Println("out:", i)
			case x <- z:
				fmt.Println("in:", z)
			}

		}(z)
	}


	// TODO: 此处会报 deadlock 错误.
	// select 不能放在 for 中，否则会造成 deadlock.
	for _,z := range []int{1,2,3,4,5,6,7,8}{

			select {
			case i := <-x:
				fmt.Println("out:", i)
			case x <- z:
				fmt.Println("in:", z)
			}
	}

	time.Sleep(time.Second * 2)
	
}
