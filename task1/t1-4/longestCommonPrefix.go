package main

import (
	"fmt"
)

/**
 * 最长公共前缀
 */
func main() {
	prefix := findLongestCommonPrefix([]string{"123", "111", "124"})
	//prefix := findLongestCommonPrefix([]string{"abc", ""})
	fmt.Println("最长公共前缀：", prefix)

}

func findLongestCommonPrefix(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}

	// 找出最短的字符串
	shortStr := strArr[0]
	for _, value := range strArr {
		if len(value) < len(shortStr) {
			shortStr = value
		}
	}
	// 若最短的字符串为空，则最长公共前缀为空
	if len(shortStr) == 0 {
		return ""
	}

	// 相同前缀
	/*var samePrefixBuilder strings.Builder
	for i := 0; i < len(shortStr); i++ {
		shortCh := shortStr[i]
		indexSameValue := true
		for _, str := range strArr {
			if str[i] != shortCh {
				indexSameValue = false
				break
			}
		}
		if indexSameValue {
			samePrefixBuilder.WriteByte(shortCh)
		} else {
			return samePrefixBuilder.String()
		}
	}
	return samePrefixBuilder.String()*/

	// 相同前缀，优化版本
	for i := range shortStr {
		for _, s := range strArr {
			if s[i] != shortStr[i] {
				return shortStr[:i]
			}
		}
	}

	return shortStr
}
