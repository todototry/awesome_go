package main

import "fmt"
import "time"
import "crypto/rand"

func main() {
	c := make(chan *[16]byte)

	go func() {
		// Use double arrays to avoid data racing.
		var dataA, dataB = new([16]byte), new([16]byte)
		for {
			_, err := rand.Read(dataA[:])
			if err != nil {
				close(c)
			} else {
				c <- dataA
				dataA, dataB = dataB, dataA
			}
		}
	}()

	for data := range c {
		fmt.Println((*data)[:])
		time.Sleep(time.Second / 2)
	}
}