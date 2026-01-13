package main

/**
 * 校验是否是回文数
 */
func main() {
	palindromeNum := 11211
	isPalindrome := validatePalindrome(palindromeNum)
	if isPalindrome {
		println("回文数")
	} else {
		println("非回文数")
	}
}

func validatePalindrome(num int) bool {
	// 负数不是回文数
	if num < 0 {
		return false
	}

	//尾数为0不是回文数
	if num%10 == 0 {
		return false
	}

	// 个数都是回文数
	if num < 10 {
		return true
	}

	originNum := num
	reverseNum := 0

	for num > 0 {
		reverseNum = reverseNum*10 + num%10
		num = num / 10
	}

	return originNum == reverseNum

}
