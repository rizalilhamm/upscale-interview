package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomOrder(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

func recursivePrint(nums []int, index int) {
	if index == len(nums) {
		return
	}

	fmt.Print(nums[index], " ")
	recursivePrint(nums, index+1)
}

func main() {
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = i + 1
	}
	nums = generateRandomOrder(nums)

	recursivePrint(nums, 0)
	fmt.Println()
}
