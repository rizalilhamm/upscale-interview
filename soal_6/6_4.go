package main

import (
	"fmt"
	"math"
)

func findMinMax(arr []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	array := []int{5, 2, 9, 1, 5, 6}
	min, max := findMinMax(array)

	fmt.Printf("Nilai terkecil: %d\n", min)
	fmt.Printf("Nilai terbesar: %d\n", max)
}
