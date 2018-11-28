package main

import (
	"fmt"
)

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func g() (result int) {
	defer func() {
		result += 1
	}()
	return 0
}

func h() (z int) {
	z = 1
	defer func(z int) {
		z += 1
	}(z)
	return z
}

func i() (r int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func j() (r int) {
	t := 5
	defer func(t int) {
		t += 5
	}(t)
	return t
}

func main() {
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(h())
	fmt.Println(i())
	fmt.Println(j())
}

// return 并非原子操作， defer 在中间, .
// 参考<go internals>
