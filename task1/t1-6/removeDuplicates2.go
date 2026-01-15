package main

import "fmt"

/**
 * 删除有序数组中的重复项 II   leetcode 标准解法
 */
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 慢指针 slow 指向下一个不重复元素应该放置的位置
	slow := 1

	// 快指针 fast 遍历数组
	for fast := 1; fast < len(nums); fast++ {
		// 如果当前元素与前一个元素不同
		if nums[fast] != nums[fast-1] {
			// 将当前元素复制到慢指针位置
			nums[slow] = nums[fast]
			// 慢指针前进
			slow++
		}
	}

	// 注意：这里没有返回新切片，只是修改了原切片的部分元素
	// 调用者应该使用 nums[:slow] 获取去重后的部分

	return slow
}

func main() {
	nums := []int{1, 2, 2, 2, 2, 2, 2, 9}

	length := removeDuplicates2(nums)

	fmt.Println("去重后长度:", length)
	fmt.Println("原数组前length个元素:", nums[:length])
	fmt.Println("整个数组:", nums)
	// 注意：后面还有重复元素，但通常我们只关心 nums[:length]
}
