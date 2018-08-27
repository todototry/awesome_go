package main

import "time"

// 1. goroutine 会立即计算、 复制执行参数
var c int // 默认值为0
func counter() int {
	c++
	return c
}

func goRoutineParam() {
	a := 100
	go func(x,y int) {
		time.Sleep(time.Second)
		println(x,y)
	}(a, counter())

	a += 100
	println(a, counter())
	time.Sleep(time.Second * 2)
	}


// 2. wait (单协程退出): 进程退出时，不会等待协程退出


// 3. waitGroup （多协程退出）： sync.WaitGroup


// 4. 暂停： runtime.Gosched()

// 5. 停止当前任务： runtime.Goexit()

// 6. 通道：


// 7. 收发：ok-idiom 或 range 模式



// 8. 单向： var send chan<- int = make(chan int)



// 9. 选择： select { }


// 10. 模式



// 11. 性能 ： 数据打包、减少加锁，提升性能



// 12. 资源泄露： 因 chan 等待造成的资源泄露。



// 13. 同步：通道并不是用来取代锁的，chan倾向于解决逻辑层次的并发处理架构，锁则用来保护局部范围的数据安全



func main(){
	// 1. goroutine 会立即复制执行参数， 然后启动计算
	goRoutineParam()

}
