package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var x = []int{1, 3, 4}
	var y = []string{"abc", "xyz"}

	fmt.Println(x, y)
	fmt.Println("swap pointer of x and y")

	xp := (unsafe.Pointer(&x))
	yp := (unsafe.Pointer(&y))
	fmt.Printf("%#v  %#v \n", *(*[]int)(xp), *(*[]string)(yp))
	xp, yp = yp, xp
	fmt.Printf("%#v  %#v \n", *(*[]string)(xp), *(*[]int)(yp))

	// 不同类型的slice pointer 想要转换，需要使用 reflect.SliceHeader 做转换。
	// slice pointer convert by reflect.sliceHeader
	// reflect.SliceHeader
	xsp := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	ysp := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	fmt.Printf("%v  %v   %v  %v\n", xsp, ysp, (*[]int)(unsafe.Pointer(xsp)), (*[]string)(unsafe.Pointer(ysp)))
	xsp, ysp = ysp, xsp
	fmt.Printf("%v  %v   %v  %v\n", xsp, ysp, (*[]string)(unsafe.Pointer(xsp)), (*[]int)(unsafe.Pointer(ysp)))

	// reflect.StringHeader
	// reflect.MakeFunc()
}
