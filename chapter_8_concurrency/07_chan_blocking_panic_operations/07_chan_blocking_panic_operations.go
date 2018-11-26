package main

func main() {
	// TODO: all the following stmt will cause deadlock.

	// 1. Following ways can be used to block current goroutine ‎forever:
	// receive from a channel which no values will be sent to
	<-make(chan struct{})
	// or
	<-make(<-chan struct{})

	// 2. send value to a channel which no ones will receive values from
	make(chan struct{}) <- struct{}{}
	// or
	make(chan<- struct{}) <- struct{}{}
	// 3. receive value from a nil channel
	<-chan struct{}(nil)
	// or
	for range chan struct{}(nil) {}
	// 4. send value to a nil channel
	(chan struct{})(nil) <- struct{}{}
	// 5. use a bare select block
	select{}
}
