package main

import (
	"fmt"
	"time"
)

func mySleep(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}

func main() {
	fmt.Println("start")
	mySleep(30)
	fmt.Println("finish")
}
