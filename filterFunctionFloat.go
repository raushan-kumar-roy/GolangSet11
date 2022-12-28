package main

import (
	"fmt"
)

func Filter(slice []float64, f func(float64) bool) []float64 {
	filtered := make([]float64, 0)
	for _, v := range slice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	numbers := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}

	noDecimals := Filter(numbers, func(v float64) bool {
		return v == float64(int(v))
	})
	fmt.Println("Numbers which do not have any digits after the decimal:")
	fmt.Println(noDecimals)

	divisibleByThree := Filter(numbers, func(v float64) bool {
		return v/3.0 == float64(int(v/3.0))
	})
	fmt.Println("Numbers which are divisible by 3.0:")
	fmt.Println(divisibleByThree)
}
