package main

import "fmt"

//https://stackoverflow.com/questions/43073681/golang-binary-search

var arr = []int{1, 3, 4, 6, 7, 8, 10, 13, 14, 18, 19, 21, 24, 37, 40, 45, 71}

func binarySearch(arr []int, key int) (result int, searchCount int) {
	mid := len(arr) / 2
	switch {
	case len(arr) == 0:
		result = -1
	case arr[mid] > key: //значит искомый элемент левее
		result, searchCount = binarySearch(arr[:mid], key) //ищем слева
	case arr[mid] < key: //значит искомый элемент правее
		result, searchCount = binarySearch(arr[mid+1:], key) //ищем справа
	default: //arr[mid] == key
		result = mid
	}
	searchCount++
	return
}

func main() {
	index, searchCount := binarySearch(arr, 7)
	fmt.Println(index, searchCount)
}
