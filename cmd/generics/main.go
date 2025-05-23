package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func sumMapValues[K comparable, V Number](kvMap map[K]V) V {
	total := V(0)
	for _, v := range kvMap {
		total += v
	}
	return total
}

func main() {
	var fruits = map[string]int{"apple": 60, "banana": 10, "mango": 50}
	var techStocks = map[string]float64{"aapl": 196.70, "goog": 169.85, "nvda": 130.41}
	fmt.Printf("Sum of all fruit prices = %v\n", sumMapValues(fruits))
	fmt.Printf("Sum of all tech stocks is %v\n", sumMapValues(techStocks))
}
