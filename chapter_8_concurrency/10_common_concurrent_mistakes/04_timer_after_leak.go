package main

import (
	"fmt"
	"time"
)

// The function will return if a message arrival interval
// is larger than one minute.
func longRunning(messages <-chan string) {
	for {
		select {
		// 会产生大量 chan 滞留在内存
		case <-time.After(time.Minute):
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}

func main() {
	var x = make(chan string, 1)
	x <- "hello, world, chan"
	longRunning(x)

}
