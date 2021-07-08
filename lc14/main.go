package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"aa", "a"}
	fmt.Println(longestCommonPrefixTwo(strs))
}

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	// 定义一个rune类型的变量
	prefix := map[int]rune{}
	var newSlice []int
	var prefixString string
	if len(strs) > 0 {
	fix:
		for i := 0; i < len(strs[0]); i++ {
			for _, val := range strs {
				if i >= len(val) {
					delete(prefix, i)
					break fix
				}
				char := []rune(val)[i]
				if _, ok := prefix[i]; ok && char != prefix[i] {
					delete(prefix, i)
					break fix
				}
				prefix[i] = char
			}
		}
		if len(prefix) > 0 {
			for key := range prefix {
				newSlice = append(newSlice, key)
			}
			sort.Ints(newSlice)
			for _, val := range newSlice {
				prefixString += string(prefix[val])
			}
		}
	}
	return prefixString
}

func longestCommonPrefixTwo(strs []string) string {
	var prefixString string
	var prefix string
	var shortStr string
	if len(strs) != 0 {
		for _, val := range strs {
			if len(shortStr) == 0 || len(shortStr) > len(val) {
				shortStr = val
			}
		}
	fix:
		for i := 0; i < len(shortStr); i++ {
			prefix = shortStr[0 : i+1]
			for _, val := range strs {
				if strings.HasPrefix(val, prefix) { //判断字符串是否存在前缀prefix
					prefixString = prefix
				} else {
					if len(prefix) < 2 {
						prefixString = ""
					} else {
						prefixString = prefix[:len(prefix)-1]
					}
					break fix
				}
			}
		}
	}
	return prefixString
}
