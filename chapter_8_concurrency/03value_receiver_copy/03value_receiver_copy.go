package main

import (
	"sync"
	"fmt"
)

type Counter struct {
	sync.Mutex
	n int64
}

// This method is okay.
func (c *Counter) Increase(d int64) (r int64) {
	c.Lock()
	c.n += d
	r = c.n
	c.Unlock()
	return
}

// The method is bad. When it is called, a Counter
// receiver value will be copied.
func (c Counter) Value() (r int64) {
	c.Lock()
	c.n++
	r = c.n
	fmt.Printf("%P, %v %d \n", &c, &(c.n), c.n)  // address of c.n  here is different.
	c.Unlock()
	return
}

func main() {
	c := Counter{n:1}
	fmt.Printf("%P, %v %d \n", &c, &(c.n), c.n)
	c.Value()
	fmt.Printf("%P, %v %d \n", &c, &(c.n), c.n)
}
