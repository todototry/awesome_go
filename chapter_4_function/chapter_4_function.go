package main

import (
	"math/rand"
	"fmt"
	"errors"
)

// 1. 定义函数： 函数外部定义函数， 函数内部定义函数
func randomInt() int {
	return rand.Int()
}

// 1.1 定义函数类型
type event func(source int32, eventType string) (int32, string)

var f event = func(source int32, eventType string) (int32, string) {
	i := source
	status := eventType

	return i, status
}

// 2. 可变长参数 : 本质上是 slice
func ellipsisFunc(a ...int){
	fmt.Printf("%T : %v", a, a)
}

// 2.1 如果参数是 array, 需要将array -> slice， 再使用 slice 展开
// see body of main()

// 3. 具名返回值
func namedReturn(age int, name string) (ageNxt int, nameNxt string){

	return age+1, name + "1"
}

// 3.1 是否有默认值？ 默认值是什么
func namedReturn2(age int, name string) (ageNxt int, nameNxt string){

	return ageNxt, nameNxt
}



// 4. 匿名函数
// see the body content of main()


// 5. 闭包


// 6. 延迟调用 defer



// 7. 错误处理： error


// 8. panic , recover



func main(){
	// 0.  不支持默认参数 default-value args not supported


	// 1. 定义函数： 函数外部定义函数， 函数内部定义函数
	x := randomInt()
	println(x)

	y,z := f(1, "success")
	println(y,z)

	// 2. 可变长参数: 本质上是 slice
	ellipsisFunc(1,2,3,4,5,6,7,8,9)

	// 2.1 如果参数是 array, 需要将array -> slice， 再使用 slice 展开
	a := [3]int{1,2,3}
	ellipsisFunc(a[:]...)

	// 3. 具名返回值:
	ageN, nameN := namedReturn(11, "fdy")
	println(ageN, nameN)

	// 3.1 是否有默认值？ 默认值是什么
	ageN2, nameN2 := namedReturn2(11, "fdy")
	fmt.Printf("\n 具名函数的默认返回值：\n %T, %v, \n %T, %v \n\n\n", ageN2, ageN2, nameN2, nameN2)


	// 4. 匿名函数
	noNameFunc := func(age int, name string) (int, string) {

		return age, name
	}
	age, name := noNameFunc(11, "fdy")
	println(age, name)

	// 5. 闭包：函数 + 上下文环境的指针引用
	// 5.1 闭包 需要注意指针的延迟计算问题，造成的值一致问题。
	closure := func(context int) func() {
		return func(){
			println("data in closure from outside: ", context)
		}
	}
	resFunc := closure(9)
	resFunc()

	// 6. 延迟调用 defer
	// defer 比手动调用，需要额外的注册、调用的性能开销
	// defer 遵循 FILO 原则
	userDefer := func() {
		defer func() {
			println("register defer 1st")
		}()

		defer func() {
			println("register defer 2nd")
		}()

		defer func() {
			println("register defer 3rd")
		}()
	}

	userDefer()



	// 7. 错误处理： error
	var errDivByZero = errors.New("division by zero.")

	div := func (a int, b int) (int, error){
		if b == 0 {
			return 0, errDivByZero
		}
		return a/b, nil
	}

	res, err := div(0, 1)
	res2, err2 := div(1, 0)
	println(res, err)
	println(res2, err2)

	// 8. panic , recover
	// 8.1 recover 只能在 Defer 延迟调用函数中执行才能正常工作
	// 8.2 panic 和 recover 必须连续地成对出现，才能多个
	// 8.3 综合以上，如果要像 try-catch 一样保护一段代码，必须要将代码放入一段匿名函数中，如下：

	divSafe := func(a int, b int) (z int, err error) {
		err = nil
		z = 0

		// 1. 将要保护的代码片段 z=a/b 放入 匿名函数， 并在其中加入 defer + recover 保护
		func() {

			// recover 必须放入在 defer 中
			defer func() {
				if recover() != nil {
					z = 0
					err = errDivByZero
				}
			}()

			z = a/b
		}()

		return z, err
	}

	res3, err3 := divSafe(1,0)
	if err3 == nil {
		println("divSafe with panic and recover: ", res3,  err3)
	} else {
		println("divSafe with panic and recover: ", res3,  err3.Error())
	}

	res4, err4 := divSafe(1,1)
	if err4 == nil {
		println("divSafe with panic and recover: ", res4,  err4)
	} else {
		println("divSafe with panic and recover: ", res4,  err4.Error())
	}

}