package main

import (
	"fmt"
	"reflect"
)

func getType(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func main() {
	a := 5
	b := "test"
	c := true
	d := make(chan int)

	getType(a)
	getType(b)
	getType(c)
	getType(d)
}
