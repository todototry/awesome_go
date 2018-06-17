package main

import (
	"fmt"
)

/* 9大复杂类型： 参考源码 type.go 底部
func (t *Basic) Underlying() Type     { return t }
func (t *Array) Underlying() Type     { return t }
func (t *Slice) Underlying() Type     { return t }
func (t *Struct) Underlying() Type    { return t }
func (t *Pointer) Underlying() Type   { return t }
func (t *Tuple) Underlying() Type     { return t }
func (t *Signature) Underlying() Type { return t }
func (t *Interface) Underlying() Type { return t }
func (t *Map) Underlying() Type       { return t }
func (t *Chan) Underlying() Type      { return t }
func (t *Named) Underlying() Type     { return t.underlying }
*/

func main() {
	// map string : int
	x := make(map [string] int)
	x["x"]=1
	fmt.Println(x,"Map")
	result, ok := x["x"]
	fmt.Println(result, ok, "Result, ok := x['x']")

	res, ok := x["y"]
	fmt.Println(res, ok, "res, ok := x['y']")

	// arr


	// slice
	y := make([]int, 0, 5)
	print(y, "Slice")

	// chan


	}
