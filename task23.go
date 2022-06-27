package main

import "fmt"

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...) //сдвигаем все элементы справа от удаляемого на один влево
}

func main() {
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	fmt.Println(slice)
	out := remove(slice, 5)
	fmt.Println(out)
}
