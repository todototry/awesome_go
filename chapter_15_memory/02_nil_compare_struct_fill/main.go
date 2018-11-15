package main

import (
	"fmt"
	"unsafe"
)

// http://127.0.0.1:55555/article/summaries.html#nil-zero-value-types

// TODO: 变量自动设置为 nil 的初始化情形
func new_nil() {
	fmt.Println("--------ptr new------------")
	fmt.Println(*new(*int) == nil)         // true
	fmt.Println(*new([]int) == nil)        // true
	fmt.Println(*new(map[int]bool) == nil) // true
	fmt.Println(*new(chan string) == nil)  // true
	fmt.Println(*new(func()) == nil)       // true
	fmt.Println(*new(interface{}) == nil)  // true

	fmt.Println("---------new-----------")
	fmt.Println(new(*int) == nil)         // false
	fmt.Println(new([]int) == nil)        // false
	fmt.Println(new(map[int]bool) == nil) // false
	fmt.Println(new(chan string) == nil)  // false
	fmt.Println(new(func()) == nil)       // false
	fmt.Println(new(interface{}) == nil)  // false

	fmt.Println("---------var x-----------")

	var aintptr *int
	var intslice []int
	var mapint map[int]bool
	var chanstr chan string
	var funcptr func()
	var object interface{}

	fmt.Println(aintptr == nil)         // false
	fmt.Println(intslice == nil)        // false
	fmt.Println(mapint == nil) // false
	fmt.Println(chanstr == nil)  // false
	fmt.Println(funcptr == nil)       // false
	fmt.Println(object == nil)  // false

	fmt.Println("---------make and var len()-----------")


	// TODO: 特别点, nil
	var xchanint chan int
	var xslice []int
	var xmapint map[string]int
	fmt.Println( xchanint == nil)                     // true
	fmt.Println( xslice == nil)                       // true
	fmt.Println( xmapint == nil)                      // true

	// false
	var xxmakechan = make(chan int)
	var xxmakechan1 = make(chan int, 1)
	var xxmakechan2 = make(chan int, 2)
	var makeslice = make(chan int)
	var makeslice1 = make(chan int, 1)
	var makeslice2 = make(chan int, 2)
	var makemap = make(map[string]int)
	var makemap1 = make(map[string]int, 1)
	var makemap2 = make(map[string]int, 2)

	fmt.Println( xxmakechan == nil, len(xxmakechan), cap(xxmakechan))  // false
	fmt.Println( xxmakechan1 == nil, len(xxmakechan1), cap(xxmakechan1))  // false
	fmt.Println( xxmakechan2 == nil, len(xxmakechan2), cap(xxmakechan2))  // false
	fmt.Println( makeslice == nil, len(makeslice), cap(makeslice))  // false
	fmt.Println( makeslice1 == nil, len(makeslice1), cap(makeslice1))  // false
	fmt.Println( makeslice2 == nil, len(makeslice2), cap(makeslice2))  // false
	fmt.Println( makemap == nil, len(makemap))  // false
	fmt.Println( makemap1 == nil, len(makemap1))  // false
	fmt.Println( makemap2 == nil, len(makemap2))  // false
	fmt.Println("--------------------")

}

func quick_fill_nil()  {
	//哪些情形下初始化的 value 是 nil?

	// This following line doesn't compile.
	// v := nil

	// There must be sufficient information for compiler
	// to deduce the type of a nil value.
	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// This lines are equivalent to the above lines.
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

}

func quick_fil_zero(){
	// 空结构体 长度是0，指向的是 runtime.zerobase
	var x = struct {

	}{}
	fmt.Println(x)

	// 空结构体 内存占用的长度是0，指向的是 runtime.zerobase ,
	var s  [10]struct{}
	// len 和 cap 都是10  ， 但是 size 是0 ，屌
	fmt.Println(s, len(s), cap(s), unsafe.Sizeof(s))

	// 元素默认初始化为0
	var y = struct {
		a int
	}{}

	fmt.Println(y.a)


}

func nil_compare(){
	//1. In Go, map. slice and function types don't support comparison(不支持比较). So comparing two nil identifiers specified with any type of the uncomparable types is illegal.
	// The following lines fail to compile.
	//var _ = ([]int)(nil) == ([]int)(nil)
	//var _ = (map[string]int)(nil) == (map[string]int)(nil)
	//var _ = (func())(nil) == (func())(nil)


	//2. But any values of the above mentioned uncomparable types can be compared with a bare nil identifier.
	// The following lines compile okay.
	var _ = ([]int)(nil) == nil
	var _ = (map[string]int)(nil) == nil
	var _ = (func())(nil) == nil

}

func complex_size() {
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

}

func main() {
	new_nil()
	quick_fill_nil()
	quick_fil_zero()
	nil_compare()
	complex_size()
	}
