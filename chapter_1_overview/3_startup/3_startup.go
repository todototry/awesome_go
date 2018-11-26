// 0. 全局： 按照 import 顺序，逐个初始化 import 的 pkg
// 1. import pkg内部，先初始化const， var， 然后是 init, 可以有多个 init 函数，逐个按照顺序初始化
// 2. 最后回到main pkg， 初始化 const， var， init，最后执行 main

// -------注意--------：
// 以上所有步骤都在主 goroutine 中执行，如果在某个 init 中使用了 go func() ， 使用了新的 go routine, 只能在 main.main（）启动后被执行。
// -------注意--------：

package main

import (
	"awson_go/chapter_1_overview/3_startup/subpkg"
	"fmt"
)

func init() {
	fmt.Println(" 1 in main.")
}

func init() {
	fmt.Println(" 2 in main.")
}

func main() {
	subpkg.Dep()
	x := make(chan int, 1)

	x <- 1

	go func() {
		<-x
		<-x
	}()

	<-x
	
}
