package main

import (
	"fmt"
	"sync"
)

var a [5]int = [5]int{2, 4, 6, 8, 10}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(a[i] * a[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
}
