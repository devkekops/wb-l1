package main

import (
	"fmt"
	"time"
)

func worker(in chan int, out chan int) {
	for {
		num, ok := <-in
		if !ok {
			close(out)
			return
		}
		out <- num * 2 //для завершения главного потока, который вычитывает из out
	}
}

func main() {
	arr := []int{2, 4, 78, 4, 9, 11, 23, 123, 10}

	in := make(chan int)
	out := make(chan int)
	go worker(in, out)

	go func() {
		for i := 0; i < len(arr); i++ {
			in <- arr[i]
			fmt.Printf("sent: %d\n", arr[i])
			time.Sleep(time.Second)
		}
		close(in) //для завершения worker-а
	}()

	for {
		num, ok := <-out
		if !ok {
			break
		}
		fmt.Printf("received: %d\n", num)
	}
}
