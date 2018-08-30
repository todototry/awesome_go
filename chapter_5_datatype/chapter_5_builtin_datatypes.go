package main

import (
	"fmt"
	"unsafe"
	"reflect"
	"unicode/utf8"
)

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
	// 1.1 不转义声明
	s1 := ` 
		Line 1 \r\n 
		Line2 ...
		line 3. \%\r
		`
		println(s1)
	// 1.2 切片与地址不变性
	s2 := s1[:25]
	fmt.Printf("对象地址： s1: %s, s2: %s\n", &s1, &s2)
	// unsafe.Pointer 用于指针类型转换
	fmt.Printf("对象地址： s1: %s, s2: %s \n", unsafe.Pointer(&s1), unsafe.Pointer(&s2))

	// reflect.StringHeader 和 string 头结构相同
	fmt.Printf("真实字符数据的内存地址： s1: %s, s2: %s \n", (*reflect.StringHeader)(unsafe.Pointer(&s1)), (* reflect.StringHeader)(unsafe.Pointer(&s2)))
	// reflect.String
	// 1.3.1 for  遍历
	for i:=0; i<len(s1); i++ {
		fmt.Printf("%c",s1[i])
	}
	fmt.Println("完成 S1")


	for i:=0; i<len(s2); i++ {
		fmt.Printf("%c",s2[i])
	}
	fmt.Println("完成 S2")

	// 1.3.2 for range 遍历
	for index, ele := range s1{
		fmt.Printf("%v : %c \n", index, ele)
	}
	fmt.Println("完成 range")

	// 1.3.3 类型转换: []byte  ,  []rune 区别: 只有内存长度的区别，byte 是 uint8 , rune 是 uint32
	var bs  = []byte(s1)
	fmt.Printf("%c", bs[42])

	var br  = []rune(s1)
	fmt.Printf("%c", br[42])

	fmt.Println(cap(bs))
	fmt.Println(cap(br))

	fmt.Printf("[]byte 初始地址：%p  , 第2元素地址 %p , 差值：  \n", &bs[0], &bs[1])
	fmt.Printf("[]rune 初始地址：%p  , 第2元素地址 %p ，差值：  \n", &br[0], &br[1])

	// 1.4 操作性能


	// 1.5 Unicode: rune 本质上是专门用来存储 Unicode 码点的


	// 1.6 字符变量与字符串长度：
	ChineseStr := "中.文"
	fmt.Printf("len(): %v , utf8.RuneCount %v \n", len(ChineseStr), utf8.RuneCountInString(ChineseStr))

	// 2. array ：初始化方法、默认值、缺省长度声明、展开操作、多维声明、指针数组和数组指针、值类型导致的复制、切片转换方法

	const (
		Invalid  = iota
		Bool
		Int
		Int8
		Int16
		Int32
		Int64
		Uint
		Uint8
		Uint16
		Uint32
		Uint64
		Uintptr
		Float32
		Float64
		Complex64
		Complex128
		Array
		Chan
		Func
		Interface
		Map
		Ptr
		Slice
		String
		Struct
		UnsafePointer
	)
	fmt.Printf("%v", Int)


	var kindNames = []string{
		Invalid:       "invalid",
		Bool:          "bool",
		Int:           "int",
		Int8:          "int8",
		Int16:         "int16",
		Int32:         "int32",
		Int64:         "int64",
		Uint:          "uint",
		Uint8:         "uint8",
		Uint16:        "uint16",
		Uint32:        "uint32",
		Uint64:        "uint64",
		Uintptr:       "uintptr",
		Float32:       "float32",
		Float64:       "float64",
		Complex64:     "complex64",
		Complex128:    "complex128",
		Array:         "array",
		Chan:          "chan",
		Func:          "func",
		Interface:     "interface",
		Map:           "map",
		Ptr:           "ptr",
		Slice:         "slice",
		String:        "string",
		Struct:        "struct",
		UnsafePointer: "unsafe.Pointer",
	}

	fmt.Println(kindNames[0])

	var x bool = false

	fmt.Println(x)


	// 2.2 数组指针 与 指针数组



	// 3. slice：切片是数组的包装，是只读对象、对应数组必须 addressable、hash 表中数组不是 addressable、make创建方法、len / cap 计算及含义

	// 3.1 数组与切片的创建区别 ？
	var arr [2]int // 数组的创建，默认值为0
	var slc = []int{} //创建 slice
	var nilSlice []int  //创建一个nil slice
	fmt.Printf("%v ,\n %v, \n %v \n，%p \n，%p \n， %p \n，", arr, slc, nilSlice, arr, slc, nilSlice)

	// len , cap
	s3 := []int{10,20, 4:30}  //  :号前 为 index, 后为 value ,中间补充默认值0
	fmt.Printf("%v , len(): %v , cap(): %v \n", s3, len(s3), cap(s3))

	// 3.2 make 的开销、 reslice 、 append 、 copy
	var stack = make([]int, 3, 5)
	fmt.Printf("%p , %v ,%v  , %v \n", stack, stack, len(stack), cap(stack))

	// 3.3 切片自动增长原则：若使用 append 增加数据时，超出 cap 限制，会造成自动创建一个新的切片；
	fmt.Println("3.3 切片自动增长原则：若使用 append 增加数据时，超出 cap 限制，会造成自动创建一个新的切片；意味着任何修改都不会影响原数组，而是在新内存块上")
	var slca = make([]int , 3, 5)
	slc_append_overflow := slca[:2:4]
	fmt.Printf("old: %p , %v , len: %v  , cap: %v \n ", slc_append_overflow, slc_append_overflow, len(slc_append_overflow), cap(slc_append_overflow))
	slc_append_overflow = append(slc_append_overflow, 1,2,3)
	fmt.Printf("new: %p , %v , len: %v  , cap: %v \n ", slc_append_overflow, slc_append_overflow, len(slc_append_overflow), cap(slc_append_overflow))
	fmt.Printf("origin: %p , %v , len: %v  , cap: %v \n ", slca, slca, len(slca), cap(slca))

	// 3.4 使用 append 增加数据时， 超出 len() 但是未超出 cap 是否创建新切片呢 ？
	fmt.Printf("3.4 使用 append 增加数据时， 超出 len() 但是未超出 cap 是否创建新切片, 还是会沿用原数组， 意味着任何修改都会影响原数组")
	slc_append := slca[:2:4]
	fmt.Printf("origin: %p , %v , len: %v  , cap: %v \n ", slca, slca, len(slca), cap(slca))
	fmt.Printf("old: %p , %v , len: %v  , cap: %v \n ", slc_append, slc_append, len(slc_append), cap(slc_append))
	slc_append = append(slc_append, 1,)
	fmt.Printf("new: %p , %v , len: %v  , cap: %v \n ", slc_append, slc_append, len(slc_append), cap(slc_append))
	fmt.Printf("origin: %p , %v , len: %v  , cap: %v \n ", slca, slca, len(slca), cap(slca))


	// 4. map ： 创建方法、key 的要求、默认值、ok-idiom、每次迭代顺序不确定、hash 造成的 not addressable、 更新方法、读写安全性、性能
	var mapint = map[string] int {}
	mapint["X"] = 1

	type User struct{
		name string
		age uint8
	}

	// map 的更新： 由于字典被设计成 not address ,因此直接取元素内部的成员会引起赋值会失败
	var userdict = make(map[int] User)
	userdict[1] = User{
		name:"fdy",
		age:10,
	}
	// 语法出现错误， not addressable userdict[1].age += 2
	u := userdict[1]
	u.age += 3
	userdict[1] = u
	fmt.Printf("%v", userdict)


	// 5. struct：初始化（顺序法、命名法）、空结构、匿名字段、字段标签、内存布局/内存对齐
	

	// 6. Pointer.


	// 7. Tuple.


	}
