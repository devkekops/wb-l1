package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("218882428714186575617", 0)
	b, _ := new(big.Int).SetString("111111111111111111111", 0)

	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Sub(a, b))
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Div(a, b))
}
