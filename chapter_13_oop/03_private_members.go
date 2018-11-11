package main

import (
	"encoding/json"
	"fmt"
)

type personp struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := personp{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
