package main

import "fmt"

type N int

func (n N) value() N {
	n++  										   // 复制后以第一个参数传入value()函数中， 因此内部修改不影响外部的n
	fmt.Println("Origin++: %p, %v", &n ,n)
	return n
}

func (n *N) pointer() N {
	(*n)++
	fmt.Println("Pointer++: %p, %v", &n, *n)
	return *n
}


func main(){
	var a N = 25

	fmt.Printf("N: %p, %v \n", &a, a)
	p := &a

	a.value()
	p.value()

	a.pointer()
	p.pointer()

	p.value()

	a.pointer()
	p.pointer()
}
