package main

import "fmt"

// 1. 定义接口、接口默认值


// 2. 实现检测、匿名接口

type itester interface {
	test()
	string() string
}

type data struct {

}

func (*data) test(){

}

func (data) string() string {
	s := "hi fmt.String"
	fmt.Println(s)
	return s
}



// 3. 超集接口、 子集接口


// 4. 执行机制：
type Ner interface {
	a()
	b(int)
	c(string) string
}
type N int
func (*N) a(){}
func (*N) b(int) {}
func (*N) c(s string) string {
	return s
}

// 5. 类型转换： 空接口 + ok-idiom的接口转换 和 接口实现检测


// 6. 技巧



func main(){
	// demo 1

	// demo 2:
	var d data
	// var t itester = d data 中有一个函数的 receiver 是 *data
	var t itester = &d

	t.test()
	t.string()

	// 接口临时类型转换
	t.(*data).string()

	// 接口类型转换， 使用 ok-idiom 模式，即使失败也不会引发 panic
	if s, ok := t.(*data); ok {
		fmt.Println("-----convert back to data----")
		fmt.Printf("%T \n", s)
		s.string()
	}

	// 接口实现检测： type switch
	switch v := t.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println("fmt.Stringer")
	case *data:
		fmt.Println(v)
		fmt.Println("*data")
	default:
		fmt.Println("default")

	}



	// demo 4:
	var n N
	var nn Ner = &n
	fmt.Println(nn.c("hello world"))
}
