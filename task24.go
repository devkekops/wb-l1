package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func getDistance(a *Point, b *Point) float64 {
	xSq, ySq := math.Pow(b.x-a.x, 2), math.Pow(b.y-a.y, 2)
	return math.Sqrt(float64(xSq + ySq)) //расстояние между двумя точками - корень из (x_2 - x_1)^2 + (y_2 - y_1)^2
}

func main() {
	a := &Point{0, 0}
	b := &Point{1, 1}
	fmt.Println(getDistance(a, b))
}
