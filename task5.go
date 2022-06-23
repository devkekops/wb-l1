package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 122))
	}

	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var t int
	flag.IntVar(&t, "t", 5, "time exit")
	flag.Parse()
	var wg sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)

	ourChan := make(chan interface{})

	wg.Add(1)
	go func() {
		for {
			select {
			case data := <-ourChan:
				fmt.Printf("Received from channel: %v\n", data)
			case <-ctx.Done():
				fmt.Println("Stop goroutine")
				wg.Done()
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			os.Exit(0)
		default:
			num := rand.Intn(1000)
			str := randomString(12)
			fmt.Printf("Sent to channel: %d\n", num)
			fmt.Printf("Sent to channel: %s\n", str)

			ourChan <- num
			ourChan <- str

			time.Sleep(2 * time.Second)
		}
	}
}
