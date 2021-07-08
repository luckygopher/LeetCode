package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstringThree("我们哈哈哈"))
}

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 方法1
func lengthOfLongestSubstring(s string) int {
	stringOne := ""
	stringTwo := ""
	for _, v := range s {
		strIndex := strings.Index(stringOne, string(v))
		if strIndex == -1 {
			stringOne += string(v)
		} else {
			if len(stringOne) > len(stringTwo) {
				stringTwo = stringOne
			}
			stringOne = stringOne[strIndex+1:] + string(v)
		}
	}
	strlen := int(math.Max(float64(len(stringOne)), float64(len(stringTwo))))
	return strlen
}

//方法2
func lengthOfLongestSubstringTwo(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

//方法三 支持中文
func lengthOfLongestSubstringThree(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
