# 通道引发的资源泄露检测方法
## 原因
 因 goroutine 卡在了发送或接收操作，一直不被唤醒，垃圾回收器也不敢回收，导致一直处于等待队列里长久休眠.

## 检测方法
 1. 编译输出

    `go build -o test`
 2. go debug 环境设置 + 运行

    `GODEBUG="gctrace=1,schedtrace=1000,scheddetail=1"  ./test`

 可以看到输出中 goroutine 的 chan 处于 receive 或者 send 中的状态

## Goroutine泄露
    Go语言是带内存自动回收的特性，因此内存一般不会泄漏。但是Goroutine确存在泄漏的情况，同时泄漏的Goroutine引用的内存同样无法被回收。
    func main() {
        ch := func() <-chan int {
            ch := make(chan int)
            go func() {
                for i := 0; ; i++ {
                    ch <- i
                }
            } ()
            return ch
        }()

        for v := range ch {
            fmt.Println(v)
            if v == 5 {
                break
            }
        }
    }

上面的程序中后台Goroutine向管道输入自然数序列，main函数中输出序列。但是当break跳出for循环的时候，后台Goroutine就处于无法被回收的状态了。

## 解决方案
通过context包来避免这个问题

    func main() {
        ctx, cancel := context.WithCancel(context.Background())

        ch := func(ctx context.Context) <-chan int {
            ch := make(chan int)
            go func() {
                for i := 0; ; i++ {
                    select {
                    case <- ctx.Done():
                        return
                    case ch <- i:
                    }
                }
            } ()
            return ch
        }(ctx)

        for v := range ch {
            fmt.Println(v)
            if v == 5 {
                cancel()
                break
            }
        }
    }

当main函数在break跳出循环时，通过调用cancel()来通知后台Goroutine退出，这样就避免了Goroutine的泄漏。

