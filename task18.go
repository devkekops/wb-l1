package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter   int
	counterCh chan int
}

func (c *Counter) count() {
	for {
		i := <-c.counterCh
		c.counter = c.counter + i
	}
}

func (c *Counter) incr(i int) {
	c.counterCh <- i
}

func main() {
	cntr := Counter{0, make(chan int)}
	go cntr.count()

	gorNum := 15
	var wg sync.WaitGroup
	wg.Add(gorNum)

	for i := 0; i < gorNum; i++ {
		go func() {
			cntr.incr(1)
			wg.Done()
		}()
	}
	wg.Wait()

	close(cntr.counterCh)

	fmt.Println(cntr.counter)
}
