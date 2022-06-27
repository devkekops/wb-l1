package main

import (
	"fmt"
	"reflect"
)

//https://stackoverflow.com/questions/44956031/how-to-get-intersection-of-two-slice-in-golang/
//https://github.com/juliangruber/go-intersect/blob/master/intersect.go

var a = []int{1, 3, 5, 10, 4, 8}
var b = []int{2, 6, 3, 4, 11, 7}

func simpleIntersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			set = append(set, el)
		}
	}
	return set
}

func hashIntersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	hash := make(map[interface{}]bool)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = true
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}

func contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(simpleIntersect(a, b))
	fmt.Println(hashIntersect(a, b))
}
