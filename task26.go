package main

import (
	"fmt"
	"strings"
)

var in = []string{"abcd", "abCdefAaf", "aabcd"}

func checkUnique(in string) bool {
	test := []rune(strings.ToLower(in))
	check := make(map[rune]bool)
	for _, letter := range test {
		if _, ok := check[letter]; ok {
			return false
		} else {
			check[letter] = true
		}
	}
	return true
}

func main() {
	for i := 0; i < len(in); i++ {
		fmt.Println(checkUnique(in[i]))
	}
}
