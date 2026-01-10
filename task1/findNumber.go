package main

import (
	"errors"
	"fmt"
)

/**
 * 找出只出现一次的数字
 */
func main() {

	numberArr := []int{1, 1, 2, 3, 3, 3}

	number, err := findFirstNumber(numberArr)
	if err != nil {
		fmt.Println("执行异常：", err)
	} else {
		if number == -1 {
			fmt.Println("不存在只出现一次的数字")
		} else {
			fmt.Println("只出现一次的数字", number)
		}
	}

}

func findFirstNumber(numberArr []int) (int, error) {
	if len(numberArr) == 0 {
		return -1, errors.New("入参数组不能为空")
	}

	number2CountMap := make(map[int]int)
	for _, number := range numberArr {
		count, exists := number2CountMap[number]
		if exists {
			number2CountMap[number] = count + 1
		} else {
			number2CountMap[number] = 1
		}
	}

	for number, count := range number2CountMap {
		if count == 1 {
			return number, nil
		}
	}

	fmt.Println("不存在只出现一次的数字")
	return -1, nil
}
