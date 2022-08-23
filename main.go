package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	result := search(nums, target)
	fmt.Println(result)
}
func search(nums []int, target int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = i
			break
		} else {
			result = -1
		}
	}
	return result
}
