package main

import (
	"unsafe"
	"fmt"
)


func addr_change_by_GC(){
	// Go语言中对象的地址可能发生变化，因此指针不能从其它非指针类型的值生成：
	// 当内存发送变化的时候，相关的指针会同步更新，但是非指针类型的uintptr不会做同步更新。
	//
	//同理CGO中也不能保存Go对象地址。

	var x int = 42
	var p uintptr = uintptr(unsafe.Pointer(&x))

	//runtime.GC()
	var px *int = (*int)(unsafe.Pointer(p))
	fmt.Println(*px, p, px)
	fmt.Printf("address changed, from %v to %v \n", p, px)
	fmt.Printf("pointer vars' addressses:  %v  %v \n", &p, &px)
	fmt.Printf("original value: %v  %v \n", *(*int)(unsafe.Pointer(p)), *px)  // 指针类型转换
}

func ptr_convert_and_offset()  {
	var x struct {
		a bool
		b int16
		c []int
	}

	/**
	unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
	*/

	/**
	uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	指针的运算
	*/
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
	fmt.Println(uintptr(unsafe.Pointer(&x)) ,  unsafe.Offsetof(x.b), uintptr(unsafe.Pointer(&(x.b))))
}


func main() {
	addr_change_by_GC()
	ptr_convert_and_offset()
}
