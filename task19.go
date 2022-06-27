package main

import "fmt"

var src = "главрыба"

//rune — это алиас к int32. Как и байты, руны были созданы для отличия от встроенного типа данных. Каждая руна представляет собой код символа стандарта Юникод.

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(src)
	fmt.Println(reverse(src))
}
