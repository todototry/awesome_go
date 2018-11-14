package main

import (
	"fmt"
	"time"
)


var cc = make(chan int)
var x = make(chan int)

func send() {
	for i := 0; i < 100; i++ {
		cc <- i
	}
}

func receive() {
	for {
		fmt.Println(<-cc)
	}
}

func anonimous_param(){
	// TODO: 本函数会发生 data race.
	// go routine，参数名 未定义， 被隐藏丢弃

	for _,z := range []int{1,2,3,4,5,6,7,8}{
		go func( int) {
			select {
			case i := <-x:
				fmt.Println("out:", i)
			case x <- 1: // TODO: data race发生与 主线程与 go程 竞争性读写 z
				fmt.Println("in:", 1)
			}

		}(z)
	}
}

func named_param_inloop(){
	// 由于 go routine 的执行启动时间不能保证，通常会在 for 执行完成后启动，因此，如果不传入具体数字 z, 那么会使用外部的 z 变量, 也就是最后一个 z 值.
	for _,z := range []int{1,2,3,4,5,6,7,8}{
		go func(z int) {
			select {
			case i := <-x:
				fmt.Println("out:", i)
			case x <- z:
				fmt.Println("in:", z)
			}
		}(z)
	}
}

func select_case(){
	xx := make(chan int, 1)
	// TODO: 此处会报 deadlock 错误.
	// select 不能放在 for 中，否则会造成 deadlock.

		select {
			case i := <-xx:
				fmt.Println("out:", i)
			case xx <- 1:
				fmt.Println("in:", 1)
		}

		select {
			case i := <-xx:
				fmt.Println("out:", i)
			case xx <- 1:
				fmt.Println("in:", 1)
		}
}



func select_case_in_for(){

	// TODO: 如果不设置 1， 也会产生 data race , deadlock 。
	xx := make(chan int, 1)

	// TODO: 此处会报 deadlock 错误.
	// select 不能放在 for 中，否则会造成 deadlock.
	for z:= range []int{1,2,3,4,5,6,7,8}{

		select {
		case i := <-xx:
			fmt.Println("out:", i)
		case xx <- z: // TODO: z 在 go routine 和 main 的 for 中发生数据竞争
			fmt.Println("in:", 1)
		}
	}
}


func main() {
	// 不同初始化方法带来的区别。
	var chanint = make(chan int)
	var chanint1 = make(chan int, 1)

	var varchanint chan int
	fmt.Println("var x chan int :::: value: ", varchanint, "len:", len(varchanint), "cap:", cap(varchanint))
	// TODO: failed
	// varchanint <- 1
	fmt.Println("var x chan int :::: value: ", varchanint, "len:", len(varchanint), "cap:", cap(varchanint))

	fmt.Println( chanint, chanint1,  len(chanint), len(chanint1))
	fmt.Printf("before send, value: %v, %v ，  len: %d   len: %d  \n", chanint, chanint1,  len(chanint), len(chanint1))
	fmt.Printf("before send, value: %v, %v ，  cap: %d   cap: %d  \n", chanint, chanint1,  cap(chanint), cap(chanint1))

	chanint1 <- 1

	fmt.Printf("after send to chanint1, value: %v, %v ，  len: %d   len: %d  \n", chanint, chanint1,  len(chanint), len(chanint1))
	fmt.Printf("after send to chanint1, value: %v, %v ，  cap: %d   cap: %d  \n", chanint, chanint1,  cap(chanint), cap(chanint1))

	// TODO: failed. but if chanint <- 1 runs in goroutine. it's ok.
	//chanint <- 1
	//fmt.Printf("%v, %v ，   %d    %d  \n", chanint, chanint1,  len(chanint), len(chanint1))
	//fmt.Printf("%v, %v ，   %d    %d  \n", chanint, chanint1,  cap(chanint), cap(chanint1))
	

	go send()

	go receive()

	anonimous_param()

	named_param_inloop()

	select_case()
	select_case_in_for()

	time.Sleep(time.Second * 1)

}
