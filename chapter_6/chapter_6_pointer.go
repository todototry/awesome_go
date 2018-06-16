package main

import "fmt"

type N int

func (n N) value() N {
	n++
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

	print("N: %p, %v", &a, a)
	p := &a

	a.value()
	p.value()

	a.pointer()
	p.pointer()

	p.value()

	a.pointer()
	p.pointer()
}
