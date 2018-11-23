package main

// http://go101.org/article/type-embedding.html#embeddable-types

// goroutine stack limit: 1000000000-byte

type I interface {
	m()
}

type T struct {
	I
}

func main() {
	var t T
	// t.m()  // panic: runtime error: invalid memory address or nil pointer dereference
	var i = &t
	t.I = i
	i.m() // will call t.m(), then will call i.m() again, ...
}

// The example program will dead loop and stack overflow. If you have understood the above content and polymorphism, it is easy to understand why it will dead loop.
// runtime: goroutine stack exceeds 1000000000-byte limit
