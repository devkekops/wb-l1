package main

import (
	"fmt"
	"strconv"
)

var i int64 //64 бита, от -2^63 до 2^63-1

func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

func clearBit(n int64, pos uint) int64 {
	mask := int64(^(1 << pos)) //^ - заменяет 0 на 1 в маске и наоборот
	n &= mask
	return n
}

func main() {
	j := setBit(i, 63)
	fmt.Println(j)
	fmt.Println(strconv.FormatInt(j, 2))
	k := clearBit(j, 2)
	fmt.Println(k)
	fmt.Println(strconv.FormatInt(k, 2))
}
