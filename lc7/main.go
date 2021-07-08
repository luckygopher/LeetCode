package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverse(-912300999999999999))
}

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
func reverse(x int) int {
	var s string = strconv.Itoa(x) // 将数据类型转换为string
	var res string
	var sign string = "+"
	// 判断参数符号，如果为负数则去掉符号
	if x < 0 {
		s = s[strings.Index(s, "-")+1:]
		sign = "-"
	}
	// 遍历字符串并进行反转拼接
	for _, v := range s {
		res = string(v) + res
	}
	// 将字符串结果转换为int
	number, err := strconv.Atoi(sign + res)
	if err != nil {
		fmt.Println("字符串转回int失败")
	}
	// 判断结果大小是否溢出
	if float64(number) < math.Pow(-2, 31) || float64(number) > math.Pow(2, 31)-1 {
		number = 0
	}
	return number
}
