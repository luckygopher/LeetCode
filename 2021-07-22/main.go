package main

import (
	"fmt"
	"math"
)

func main() {
	list := []int{1, 4, 9, 16, 25}
	val, ok := checkList(list)
	fmt.Println(val, ok)
}

// 输入为一个整数数列(list []int)，输出为预测改数列的下一个项的值 val 和是否找到规律 ok
// 比如 输入 []int{1,3,5,7,9} ，检测出它是等差数列，输出下一项的值 11 和找到了规律 true
// 输入数列的长度小于3时，作 无规律处理。比如：[1,3]判定无规律。[1,3,5]判定为等差数列
// 不用考虑int计算的溢出问题
func checkList(list []int) (val int, ok bool) {
	// 处理数列长度小于3的情况
	if len(list) < 3 {
		return
	}
	// 等差数列
	if val, ok = equalDifferenceSeq(list); ok {
		return
	}
	// 等比数列
	if val, ok = equalRatioSeq(list); ok {
		return
	}
	// 开方等差数列
	if val, ok = squareArithmeticSeq(list); ok {
		return
	}

	return
}

// 等差:等差数列是指从第二项起，每一项与它的前一项的差等于同一个常数的一种数列
func equalDifferenceSeq(list []int) (int, bool) {
	for i := 0; i < len(list)-2; i++ {
		d1 := list[i+1] - list[i]
		d2 := list[i+2] - list[i+1]
		if d1 != d2 {
			return 0, false
		}
	}
	n := list[len(list)-1] + (list[len(list)-1] - list[len(list)-2])
	return n, true
}

// 等比:等比数列是指从第二项起，每一项与它的前一项的比值等于同一个常数的一种数列
func equalRatioSeq(list []int) (int, bool) {
	for i := 0; i < len(list)-2; i++ {
		if list[i] == 0 || list[i+1] == 0 || list[i+2] == 0 {
			return 0, false
		}
		d1 := float64(list[i+1]) / float64(list[i])
		d2 := float64(list[i+2]) / float64(list[i+1])
		if d1 != d2 || math.Remainder(d1, float64(1)) != 0 {
			return 0, false
		}
	}
	n := list[len(list)-1] * (list[len(list)-1] / list[len(list)-2])
	return n, true
}

// 开方等差数列:
// 如果一个数列可以对每一项开根号（比如负数，则认为不可以开根号，再比如3，开根号无法得到整数，
// 也认为不可以开根号）并且得到的新的整数数列是等差数列，那么他就是开方等差数列
func squareArithmeticSeq(list []int) (int, bool) {
	equalDiff := make([]int, 0, len(list))
	for i := 0; i < len(list); i++ {
		d := math.Sqrt(float64(list[i]))
		if math.IsNaN(d) || math.Remainder(d, 1) != 0 {
			return 0, false
		}
		equalDiff = append(equalDiff, int(d))
	}
	val, ok := equalDifferenceSeq(equalDiff)
	if !ok {
		return 0, false
	}
	return int(math.Pow(float64(val), 2)), true
}
