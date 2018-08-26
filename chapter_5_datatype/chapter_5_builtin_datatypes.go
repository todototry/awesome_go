package main

var x = `/* 9大复杂类型： 参考源码 type.go 底部
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
`

func main(){
	// 1. string : 默认值、不转义声明方法、切片与地址不变性、可变性与转换 []rune / []byte 、 操作性能、Unicode、字符数量与字符串长度


	// 2. array ：初始化方法、默认值、缺省长度声明、展开操作、多维声明、指针数组和数组指针、值类型导致的复制、切片转换方法



	// 3. slice：切片是数组的包装，是只读对象、对应数组必须 addressable、hash 表中数组不是 addressable、make创建方法、len / cap 计算及含义
	// make 的开销、 reslice 、 append 、 copy


	// 4. map ： 创建方法、key 的要求、默认值、ok-idiom、每次迭代顺序不确定、hash 造成的 not addressable、 更新方法、读写安全性、性能
	var i = map[string] int {}
	i["X"] = 1

	println(i["X"])
	print(x)


	// 5. struct：初始化（顺序法、命名法）、空结构、匿名字段、字段标签、内存布局/内存对齐
	


	}
