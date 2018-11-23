package main

/*
	1. 什么是 addressable value ?
	2. 如何获得指针？ 两种方法 new + &
	3. 指针的操作和限制？操作地址对应数据。限制：不同类型的指针不能 == ,  ptr++ ，不同类型的变量的 值为nil指针也是有区别的，不能用 ==
	4. 指针是一种类型？ 与 C 语言不同的是，go 的指针类型是强类型，有无数种，并且不能互转.
	5. 指针的通用化类型的转换？ unsafe.pointer 与 uintptr 互换
	6. 进阶：指针类型转换后的使用要点?  如果是间接类型（map  slice） 需要强制转换回原来的类型的指针，才能有效操作(不是那么安全的方式)。
	7. 高级：gc 对已有变量的地址的影响： 直接变量， 含有隐藏间接数据的复杂变量(indirect hidden value: map, slice , interface , function.)
	8. 高级：打破限制：移动指针，指向 struct 的不同成员变量的方法。
	9. 高级：打破限制：指向不同类型指针的转换
	10.高级：不同类型，但是内部结构定义一致的 value，如何利用指针实现转换。
	11.高级: C 的野指针问题再 go 中。

	Refer: http://go101.org/article/unsafe.html
	http://127.0.0.1:55555/article/value-part.html
	http://127.0.0.1:55555/article/pointer.html

*/

//
//

func main() {

}
