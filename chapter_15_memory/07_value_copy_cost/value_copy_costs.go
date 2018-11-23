// http://go101.org/article/value-copy-cost.html

package main

import "testing"

type S struct{ a, b, c, d, e int64 }

var sX, sY, sZ = make([]S, 1000), make([]S, 1000), make([]S, 1000)
var sumX, sumY, sumZ int64

func Benchmark_Loop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumX = 0
		for j := 0; j < len(sX); j++ {
			sumX += sX[j].a
		}
	}
}

func Benchmark_Range_OneIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumZ = 0
		for j := range sY {
			sumZ += sY[j].a
		}
	}
}

// 当迭代器导出的是 index + value时， value 每次都是 copy 出来的。因此造成性能比 迭代只导出 value要慢。
func Benchmark_Range_TwoIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumY = 0
		for _, v := range sY {
			sumY += v.a
		}
	}
}

/*
Benchmark_Loop-4                500000   3228 ns/op
Benchmark_Range_OneIterVar-4    500000   3203 ns/op
Benchmark_Range_TwoIterVars-4   200000   6616 ns/op
*/
