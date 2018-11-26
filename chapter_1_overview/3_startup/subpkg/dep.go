package subpkg

import (
	"fmt"
)

func init() {
	fmt.Println("1st init in sub pkg.")
}

func init() {
	fmt.Println("2nd init in sub pkg.")
}

func init() {
	go func() {
		fmt.Println(" new go routine in init() from sub package.")
	}()
}

func Dep() {
	fmt.Println("func from sub package.")
}
