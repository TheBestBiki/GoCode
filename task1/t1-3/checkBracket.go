package main

/**
 * 有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
 */
func main() {
	result := isValid("{[[]}")
	if result {
		println("有效括号")
	} else {
		println("无效括号")
	}
}

func isValid(s string) bool {
	// 空字符串是有效的
	if len(s) == 0 {
		return true
	}

	// 如果字符串长度是奇数，无效
	if len(s)%2 != 0 {
		return false
	}

	// 创建映射：右括号 -> 对应的左括号
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 使用切片模拟栈
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// 如果是右括号
		if matchingLeft, isRight := pairs[ch]; isRight {
			// 栈为空或栈顶不匹配
			if len(stack) == 0 || stack[len(stack)-1] != matchingLeft {
				return false
			}
			// 匹配成功，弹出栈顶
			stack = stack[:len(stack)-1]
		} else {
			// 是左括号，入栈
			stack = append(stack, ch)
		}
	}

	// 最后栈必须为空才有效
	return len(stack) == 0
}
