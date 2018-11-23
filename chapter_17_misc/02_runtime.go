package main

import "fmt"
import "unsafe"

func main() {
	type T struct{ a int }
	var t T
	fmt.Println(&t)                                 // &{0}
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t))) // c6233120a8
	println(&t)                                     // 0xc6233120a8
}
