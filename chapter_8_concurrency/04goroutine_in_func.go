package main

import (
"fmt"
"time"
)

func main() {
	fibonacci := func() chan uint64 {
		// 可确保执行在 go func() 之前
		c := make(chan uint64)

		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y { // here
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	for x, ok := <-c; ok; x, ok = <-c { // here
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}
