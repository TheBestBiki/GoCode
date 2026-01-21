package main

import "fmt"

/*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。考察点 ：指针运算、切片操作。
*/
func main() {
	numArr := []int{1, 2, 3, 4}
	changeSlice(&numArr)
	fmt.Println(numArr)
}

func changeSlice(sliceArr *[]int) {
	copySlice := *sliceArr
	for i, v := range copySlice {
		copySlice[i] = v * 2
	}
	//*sliceArr = []int{4, 5, 6}
}
