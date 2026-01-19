package main

import "fmt"

/*
两数之和：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func main() {
	nums := []int{2, 7, 11, 15}
	targetNums := twoSum(nums, 18)
	fmt.Println(targetNums)
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 || len(nums) == 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
