package main

import (
	"context"
	"fmt"
	"time"
)

//https://www.sobyte.net/post/2021-06/several-ways-to-stop-goroutine-in-golang/

func gor1(ch chan string) {
	for {
		v, ok := <-ch //если канал закрыт, ok == false
		if !ok {
			fmt.Println("gor1 finish")
			return
		}
		fmt.Printf("gor1 receive: %s\n", v)
	}
}

func gor2(quit chan struct{}) {
	for {
		select {
		case <-quit:
			fmt.Println("gor2 finish")
			return
		default:
			fmt.Println("gor2 running")
			time.Sleep(time.Second)
		}
	}
}

func gor3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //произойдёт при вызове cancel()
			fmt.Println("gor3 finish")
			return
		default:
			fmt.Println("gor3 running")
			time.Sleep(time.Second)
		}
	}

}

func main() {
	//1 способ Закрытие канала
	ch := make(chan string)
	go gor1(ch)
	ch <- "hello"
	ch <- "world"
	close(ch)
	time.Sleep(time.Second)

	//2 способ Сигнальный канал
	quit := make(chan struct{}) //Каналы с пустой структурой вместо булева типа часто используются для оповещений о событиях
	go gor2(quit)
	time.Sleep(4 * time.Second)
	quit <- struct{}{} // Закрытый канал отдаёт нулевые значения в неблокирующем режиме. Поэтому достаточно закрыть канал, чтобы оповестить о событии все горутины
	time.Sleep(time.Second)

	//3 способ Через context
	ctx, cancel := context.WithCancel(context.Background())
	go gor3(ctx)
	time.Sleep(4 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
