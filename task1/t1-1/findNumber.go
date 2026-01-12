package main

import (
	"errors"
	"fmt"
)

/**
 * 找出只出现一次的数字
 */
func main() {

	numberArr := []int{1, 2, 3, 3, 3}

	uniqueNumberArr, err := findUniqueNumber(numberArr)
	if err != nil {
		fmt.Println("执行异常：", err)
	} else {
		// 这里就算uniqueNumberArr为nil也不会报错，可以同时处理 nil 和空切片
		if len(uniqueNumberArr) == 0 {
			fmt.Println("不存在只出现一次的数字")
		} else {
			fmt.Println("只出现一次的数字有：", uniqueNumberArr)
		}
	}

}

func findUniqueNumber(numberArr []int) (uniqueNumberArr []int, err error) {
	if len(numberArr) == 0 {
		return nil, errors.New("入参数组不能为空")
	}

	number2CountMap := make(map[int]int)
	for _, number := range numberArr {
		number2CountMap[number]++
	}

	for number, count := range number2CountMap {
		if count == 1 {
			uniqueNumberArr = append(uniqueNumberArr, number)
		}
	}

	if len(uniqueNumberArr) > 0 {
		return uniqueNumberArr, nil
	}

	fmt.Println("不存在只出现一次的数字")
	return nil, nil
}
