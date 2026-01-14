package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
 */
func main() {
	intArr := []int{14, 33, 2, 1}
	result, err := addOne(intArr)
	if err != nil {
		fmt.Println("转换异常", err)
	} else {
		fmt.Println("结果：", result)
	}
}

func addOne(intArr []int) (int, error) {
	var intBuilder strings.Builder
	for _, value := range intArr {
		intBuilder.WriteString(strconv.Itoa(value))
	}

	s := intBuilder.String()
	// 将字符串转为整数
	intValue, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return intValue + 1, nil
}
