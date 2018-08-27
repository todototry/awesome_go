package main

import (
	"time"
	"sync"
	"math"
)

// 1. goroutine 会立即计算、 复制执行参数
var c int // 默认值为0
func counter() int {
	c++
	return c
}

func goRoutineParam() {
	a := 100

	// 以下是协程内部逻辑
	go func(x,y int) {
		// 等待，让主线程先运行，更新数值，以方便对比
		time.Sleep(time.Second)
		println("go 协程内部的值：x: , y: ", x,y)
	}(a, counter())

	a += 100
	println("主线程内部的值： ", a, counter())
	time.Sleep(time.Second * 2)
	}


// 2. wait : 进程退出时，默认不会等待协程退出，使用在主线程中chan receive等待，作为单协程的退出等待方式)
func goRoutineWithChanBlockWait() {
	 blockChan := make(chan int )

	 go func() {
		 defer func() {
			 blockChan <- 1
		 }()

		 time.Sleep(time.Second)
		 println("I am in go Routine...")
	 }()

	 println("I am a code snippet behind go Routine...")
	<- blockChan
	close(blockChan)
}



// 3. waitGroup （多协程退出）： sync.WaitGroup
func MultiRoutinesWaitingWithSyncWaitGroup() {
	var waitGroup sync.WaitGroup
	res := make(chan int, 100)

	for i:=0; i<9; i++ {
		// 1. 加上要等待的信号数量
		waitGroup.Add(1)
		// 2. go work!!
		go func(i int) {
			// 3. 标记完成： tag as Done after this routine.
			defer waitGroup.Done()

			// 4. actual work.
			r := int(math.Pow10(i))
			res <- r
			println("In goRoutine :", i, " data: ", r , " was sending to result Chan..:")
		}(i)
	}

	waitGroup.Wait()

	// IMPORTANT: 使用前关闭，否则会造成下方读取消息时死锁。
	close(res)

	y := 0
	for x:= range res {
		y += x
	}
	println("accumulated result :", y)
}

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
	println("Demo 1: ")
	goRoutineParam()

	// 2. wait : 进程退出时，默认不会等待协程退出，使用在主线程中chan receive等待，作为单协程的退出等待方式)
	println("Demo 2: ")
	goRoutineWithChanBlockWait()

	// 3.
	println("Demo 3: ")
	MultiRoutinesWaitingWithSyncWaitGroup()

}
