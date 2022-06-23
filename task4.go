package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

//https://gobyexample.com/worker-pools
//https://callistaenterprise.se/blogg/teknik/2019/10/05/go-worker-cancellation/

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

type Worker struct {
	id      int
	rcvChan chan (interface{})
	ctx     context.Context
	wg      *sync.WaitGroup
}

func (w *Worker) Loop() {
	for {
		select {
		case data := <-w.rcvChan:
			fmt.Printf("worker #%d data: #%v\n", w.id, data)
		case <-w.ctx.Done():
			fmt.Printf("worker #%d finish\n", w.id)
			w.wg.Done()
			return
		}
	}
}

func main() {
	var workersNum int
	flag.IntVar(&workersNum, "w", runtime.NumCPU(), "workers number")
	flag.Parse()

	wg := &sync.WaitGroup{}
	wg.Add(workersNum)

	rcvChan := make(chan (interface{}))
	ctx, cancel := context.WithCancel(context.Background())

	workers := make([]*Worker, 0, workersNum)
	for i := 0; i < workersNum; i++ {
		workers = append(workers, &Worker{i, rcvChan, ctx, wg})
	}

	for _, w := range workers {
		go w.Loop()
	}

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-termChan
		cancel()
		wg.Wait()
		os.Exit(0)
	}()

	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(1000)
		str := randomString(12)
		fmt.Printf("Sent to channel: %d\n", num)
		fmt.Printf("Sent to channel: %s\n", str)

		rcvChan <- num
		rcvChan <- str

		time.Sleep(2 * time.Second)
	}
}
