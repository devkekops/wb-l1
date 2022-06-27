package main

import "fmt"

//множество - неупорядоченный набор уникальных элементов
//https://yourbasic.org/golang/implement-set/

var arr = []string{"cat", "cat", "dog", "cat", "tree"}

func createSet(src []string) map[string]bool {
	out := make(map[string]bool)
	for _, str := range src {
		out[str] = true
	}
	return out
}

func main() {
	for key, _ := range createSet(arr) {
		fmt.Println(key)
	}
}
