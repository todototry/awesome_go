package main

/*
引用类型: map / slice / chan
值类型：普通的 int , float , string, byte, rune

TODO: 编写输出参数和原始变量的 address 和 value 的函数
1. 引用类型的value 作为参数被传递的时候，copy 的是指针， 因此依然可以修改原变量的值。 这就是为何 chan 经常作为参数传递，函数中 send， 依然可以在函数外接收到。

TODO: 特别的， go routine 在启动前就会 copy 好相应的参数。 如果没有使用参数传入， 内部引用的 value 可能的取值不确定，因为 go 启动的时间不被保证。
2. 编写 go routine 应用外部变量， 外部变量会在外部不断更新.
*/



func main() {
	
}
