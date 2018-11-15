# 通道引发的资源泄露检测方法
## 原因
 因 goroutine 卡在了发送或接收操作，一直不被唤醒，垃圾回收器也不敢回收，导致一直处于等待队列里长久休眠.

## 检测方法
 1. 编译输出

    `go build -o test`
 2. go debug 环境设置 + 运行

    `GODEBUG="gctrace=1,schedtrace=1000,scheddetail=1"  ./test`

 可以看到输出中 goroutine 的 chan 处于 receive 或者 send 中的状态
