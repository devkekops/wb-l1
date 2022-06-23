package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s\n", h.name)
}

type Action struct {
	Human
}

func main() {
	a := Action{Human{"Oleg", 23}}
	a.SayHi()
}
