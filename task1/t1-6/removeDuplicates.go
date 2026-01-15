package main

import "fmt"

/**
 * 删除有序数组中的重复项
 */
func main() {
	nums := []int{1, 1, 2}
	result := removeDuplicates(nums)
	fmt.Println(result)
}

func removeDuplicates(nums []int) []int {
	baseLen := len(nums)
	if baseLen == 0 {
		return nums
	}
	for i := baseLen - 1; i >= 1; i-- {
		if i == 1 {
			if nums[0] == nums[1] {
				nums = nums[1:]
			}
			continue
		}
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	return nums
}
