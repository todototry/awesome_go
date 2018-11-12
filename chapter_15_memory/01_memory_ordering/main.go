package main

import "runtime"

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	//
	go setup()

	for !done {
		runtime.Gosched()
	}
	println(a) // expect to print: hello, world
}

// go run -race main.go
// go run main.go
