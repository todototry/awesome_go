package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//var x int32
	x:=1
	fmt.Println(x,"hello world")
	fmt.Printf("Now you have %g problems %f .\n", math.Pow(7,2), math.Pi)

	// 原地替换型 输出
	fmt.Printf("\r 本行将被替代 ")
	time.Sleep(time.Second)
	fmt.Printf("\r 本行已更新 ")


	println("输出到 std.error, 与 fmt.Println 混用时，顺序可能不能保证")
	fmt.Println("输出到 std.output, 与 println 混用时， 顺序可能不能保证")
}

