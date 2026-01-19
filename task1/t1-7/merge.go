package main

import (
	"fmt"
	"sort"
)

/*
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
如果没有重叠，则将当前区间添加到切片中。
*/
func main() {
	testArr := [][]int{[]int{1, 3}, []int{2, 4}, []int{6, 8}}
	resultArr := merge(testArr)
	for _, arr := range resultArr {
		fmt.Println(arr)
	}

}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 使用双指针原地合并, 写入位置
	writeIndex := 0

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := intervals[writeIndex]

		if current[0] <= last[1] {
			// 合并
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 移动到下一个位置
			writeIndex++
			intervals[writeIndex] = current
		}
	}

	return intervals[:writeIndex+1]
}
