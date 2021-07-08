package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 42, 12}
	target := 3
	res := twoSum(nums, target)
	fmt.Println(res)
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
			}
		}
	}
	return result
}
