package main

import (
	"fmt"
)

func filter(arr []float64) map[int][]float64 {
	out := make(map[int][]float64)
	for i := 0; i < len(arr); i++ {
		out[int(arr[i]/10)*10] = append(out[int(arr[i]/10)*10], arr[i])
	}

	return out
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(filter(arr))
}
