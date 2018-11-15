package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Int()
	fmt.Println(x)

}
