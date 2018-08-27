package main

import (
	"time"
	"sync"
	"math"
	"runtime"
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

// 3.1 benchmark for routine groups



// 3.2 Go Routine 没有 Local Storage 时， 协程 ID， 返回值的处理方法.
func goRoutineResult(){
	var wg sync.WaitGroup
	// 匿名struct 数组的创建
	routineInfo := [5]struct{
		id int  // 与字面量不一样， 无需 逗号 结尾
		result int
	}{}

	for i:=0; i< 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			routineInfo[i].id = i
			routineInfo[i].result = i*i
		}(i)
	}

	wg.Wait()

	// 使用 struct 数组的姿势， range 与 python 的enumerate相似
	for i, res := range routineInfo{
		println("go routine 的返回值： ", i, res.id, res.result)
	}

}


// 4. 暂停： runtime.Gosched() ， 嵌套 routine 时， 第一层的defer 依然会在 nested routine go之前执行。
func routineSchedule(){
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})
	exit2 := make(chan struct{})

	go func() {
		defer func() {
			println("  I am now in defer of 1st routine.")
			close(exit)
		}()

		go func() {
			defer func() {
				println("    I am now in defer of 2nd Nested routine.")
				close(exit2)
			}()
			println("    I am now in nested routine...")
			time.Sleep(time.Second)
			println("    I am now sleeping in nested routine...")
		}()

		println("  I am in 1st rouine...")
		for i:=0; i<=10; i++ {
			if i == 0 {
				println("  I am going to give away cpu from 1st routine....")
				//runtime.Gosched()
				println("  I am now back in 1st routine after GoSched() ....")
			}
		}


		time.Sleep(2)
		println("  I slept for 2s at 1st routine..")
	}()

	println("I am now in main function ....")
	<- exit
	println("1st routine ends")
	<- exit2
	println("2nd routine ends")
}


// 5. 停止当前任务： runtime.Goexit() : 仅仅退出当前 routine， defer依然会确保执行
func GoRuntimeExit(){
	exit := make(chan struct{})
	exit2 := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("  I am now in defer of 1st routine...")


		go func (){
			defer close(exit2)
			defer println("    I am now In defer of Nested routine...")
			println("    I am now in Nested routine before GoExit()...")
			runtime.Goexit()
			println("    I am now in Nested routine after GoExit()...")
		}()

		println("  I am now in 1st routine before GoExit()...")
		runtime.Goexit()
		println("  I am now in 1st routine after GoExit()...")
	}()


	println("I am now in main routine...")
	<- exit
	println("1st routine finished...")
	// runtime.Goexit()
	<- exit2
	println("2nd routine finished...")

}


// 6. 通道：同步模式、异步模式的使用区别，sleep 与 awake的场景 、 事件通知模型实现
func chanAsnycSync() {
	//done := make(chan struct{})
	syncChan := make(chan string)
	asyncChan := make(chan string, 100)

	// syncChan <- "hello"
	// syncChan <- "Git"
	println("syncChan size: capacity:", cap(syncChan), len(syncChan))
	println("syncChan is nil : ", syncChan == nil)
	close(syncChan)

	asyncChan <- "hello"
	println("asyncChan size: capacity:", cap(asyncChan), len(asyncChan))

}


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

	// 3. waitGroup （多协程退出等待）： sync.WaitGroup
	println("Demo 3: ")
	MultiRoutinesWaitingWithSyncWaitGroup()

	// 3.2 Go Routine 没有 Local Storage 时， 协程 ID， 返回值的处理方法.
	println("Demo 3.2 fetch routine result:")
	goRoutineResult()

	// 4. 暂停： runtime.Gosched() ， 嵌套 routine 时， 第一层的defer 依然会在 nested routine go之前执行。
	println("Demo 4. ")
	routineSchedule()

	//  5. 停止当前任务： runtime.Goexit() : 仅仅退出当前 routine， defer依然会确保执行
	println("Demo 5")
	GoRuntimeExit()

	// 6.
	chanAsnycSync()

}
