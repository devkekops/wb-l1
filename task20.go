package main

import (
	"fmt"
	"strings"
)

var str = "snow dog sun"

func reverseWords(in string) string {
	words := strings.Split(in, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(str)
	fmt.Println(reverseWords(str))
}
