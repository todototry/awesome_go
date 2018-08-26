package main

import "fmt"

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
	// 1. 定义变量：var
	// 1.1 auto detect basic
	var autoDetect = 1
	var autoStr = "this is a string"

	autoVar := 1
	autoVarStr := "this is another string"

	println(autoDetect)
	fmt.Println(autoStr)
	println(autoVar)
	println(autoVarStr)

	// 1.2 specified type
	var varInt8 int8 = 100
	var varByte byte = 'c'
	var varUint8 uint8 = 'a'

	println(varInt8)
	println(varByte)
	println(varUint8)

	// 1.3 data cast
	diff := varByte - varUint8
	diff2 := varUint8 - varByte

	println(diff)
	println(diff2)

	// 2. 定义常量： const
	const conInt = 90
	println(conInt)

	// 3. 定义枚举：Enum
	// 3.1 自动生成递增值
	const (
		a = iota
		b
		c = 100
		d
		e
		f
		g = iota
		h
		i
		j
		k
	)
	println("自定义递增值：", a,b,c,d,e,f,g,h,i,j,k)

	// 3.2 user defined
	type color byte

	const (
		black color = iota
		red
		blue
		green
	)

	println("Enum color: ", black, red, blue, green)

	// 3.3 定义枚举 Enum 组合
	const (
		_, _ = iota, iota*10
		first, firstPrize
		snd, sndPrize
		third, thirdPrize
		forth, forthPrize
	)

	println("Enum rank prize:", first, firstPrize)
	println("Enum rank prize:", snd, sndPrize)
	println("Enum rank prize:", third, thirdPrize)
	println("Enum rank prize:", forth, forthPrize)

	// 4. arr


	// 5. slice
	y := make([]int, 0, 5)
	print(y, "Slice")

	// 5.1 operation on slice.


	// 5.2



	// 6. struct



	// 7. Pointer



	// 8. Tuple


	// 9. Signature



	// 10. Interface



	// 11. map string : int
	x := make(map [string] int)
	x["x"]=1
	fmt.Println(x,"Map")
	result, ok := x["x"]
	fmt.Println(result, ok, "Result, ok := x['x']")

	res, ok := x["y"]
	fmt.Println(res, ok, "res, ok := x['y']")


	// 12. chan


	// 13. Named



}
