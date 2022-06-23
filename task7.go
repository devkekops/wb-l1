package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Test struct {
	m  map[int]string
	mu sync.Mutex //можно использовать RWMutex если бы были нужны операции чтения
}

func NewTest() *Test {
	return &Test{
		m: make(map[int]string),
	}
}

func main() {
	t := NewTest()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			t.mu.Lock()
			defer t.mu.Unlock()

			t.m[i] = strconv.Itoa(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(t.m)
}
