package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var a [5]int32 = [5]int32{2, 4, 6, 8, 10}
var sum int32 = 0

func main() {
	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt32(&sum, a[i]*a[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}
