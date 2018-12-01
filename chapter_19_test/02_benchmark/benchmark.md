# benchmark

## 编写
1. 函数以 Benchmark 开头，传入 testing.B ， 内容中 for  i <= t.N 

## 启动
go test -benchmem  awson_go/chapter_19_test/02_benchmark -bench ^Benchmark_add$

