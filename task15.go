package main

import (
	"fmt"
	"math/rand"
)

var justString string

func someFunc() {
	v := createHugeString(99)
	//justString = v[:100] //здесь может быть выход за пределы слайса, если его длина меньше 100 (panic: runtime error: slice bounds out of range [:100] with length 4)
	if len(v) < 100 {
		justString = v
	} else {
		justString = v[:100]
	}
}

func createHugeString(len int) string {
	fmt.Println(len)
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
	someFunc()
	fmt.Println(justString)
}
