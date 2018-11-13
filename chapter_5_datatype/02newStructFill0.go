package main

import (
	"fmt"
	"unsafe"
)

type Pointer struct {
	x int
	y int
}

type Circle struct {
	radius float64
	center *Pointer
}

func main() {
	cc := new(Circle)
	fmt.Printf("%v, %v, %p \n", cc, cc.radius, cc.center)

	x := Circle{}
	fmt.Printf("%v, %v, %p \n", x, x.radius, x.center)

	xx := &Circle{}
	fmt.Printf("%v, %v, %p \n", x, xx.radius, xx.center)


	var p *struct{} = nil
	fmt.Println( unsafe.Sizeof( p ) ) // 8

	var s []int = nil
	fmt.Println( unsafe.Sizeof( s ) ) // 24

	var m map[int]bool = nil
	fmt.Println( unsafe.Sizeof( m ) ) // 8

	var c chan string = nil
	fmt.Println( unsafe.Sizeof( c ) ) // 8

	var f func() = nil
	fmt.Println( unsafe.Sizeof( f ) ) // 8

	var i interface{} = nil
	fmt.Println( unsafe.Sizeof( i ) ) // 16

	_ = interface{}(nil) == 1

	}
