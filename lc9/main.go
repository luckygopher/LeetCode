package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x == reverse(x) {
		return true
	}
	return false
}

// 数字反转
func reverse(x int) int {
	var s string = strconv.Itoa(x) // 将数据类型转换为string
	var res string
	// 遍历字符串并进行反转拼接
	for _, v := range s {
		res = string(v) + res
	}
	// 将字符串结果转换为int
	number, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println("字符串转回int失败")
	}
	return number
}
