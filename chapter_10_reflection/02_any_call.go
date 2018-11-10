package main

import (
	"reflect"
	"fmt"
)

type YourT1 struct {}
func (y YourT1) MethodBar() {
	//do something
	fmt.Printf("calling ")
}

type YourT2 struct {}
func (y YourT2) MethodFoo(i int, oo string) {
	//do something
	fmt.Println(i, oo)
}

func (y YourT2) lowerCaseMethod(i int, oo string) {
	//do something
	fmt.Println(i, oo)
}


func Invoke(any interface{}, name string, args... interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

func main() {
	Invoke(YourT2{}, "MethodFoo", 10, "abc")
	// lower case Method name will fail to invoke().
	Invoke(YourT2{}, "lowerCaseMethod", 10, "abc")
	Invoke(YourT1{}, "MethodBar")
}
