package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))
}

// 罗马数字转整数
func romanToInt(s string) int {
	// 定义一个基础map
	roman := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	var number int
	for i := 0; i < len(s); i++ {
		str1 := s[i : i+1]
		var str2 string
		if i+2 <= len(s) {
			str2 = s[i : i+2]
		}
		if intVal, ok := roman[str2]; ok {
			number += intVal
			i = i + 1
		} else {
			number += roman[str1]
		}
	}
	return number
}
