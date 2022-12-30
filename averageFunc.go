package main

import "fmt"

func average(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return float64(total) / float64(len(numbers))
}

func main() {
	fmt.Println(average(1, 2, 3))
	fmt.Println(average(1, 2, 3, 4, 5))
}
