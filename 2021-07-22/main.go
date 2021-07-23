package main

import (
	"fmt"
	"math"
)

func main() {
	list := []int{2, 14, 64, 202, 502, 1062, 2004}
	val, ok := CheckList(list)
	fmt.Println(val, ok)
}

// 输入为一个整数数列(list []int)，输出为预测改数列的下一个项的值 val 和是否找到规律 ok
// 比如 输入 []int{1,3,5,7,9} ，检测出它是等差数列，输出下一项的值 11 和找到了规律 true
// 输入数列的长度小于3时，作 无规律处理。比如：[1,3]判定无规律。[1,3,5]判定为等差数列
// 不用考虑int计算的溢出问题
func CheckList(list []int) (val int, ok bool) {
	// 处理数列长度小于3的情况
	if len(list) < 3 {
		return
	}
	var a1, a2, a3 []int
	// 等差数列
	a1, val, ok = equalDifferenceSeq(list)
	if ok {
		return
	}
	// 等比数列
	a2, val, ok = equalRatioSeq(list)
	if ok {
		return
	}
	// 开方等差数列
	a3, val, ok = squareArithmeticSeq(list)
	if ok {
		return
	}

	val, ok = CheckList(a1)
	if ok {
		val += list[len(list)-1]
		return
	}

	val, ok = CheckList(a2)
	if ok {
		val *= list[len(list)-1]
		return
	}

	val, ok = CheckList(a3)
	if ok {
		val = int(math.Pow(float64(val+list[len(list)-1]), 2))
		return
	}
	return
}

// 等差:等差数列是指从第二项起，每一项与它的前一项的差等于同一个常数的一种数列
func equalDifferenceSeq(list []int) ([]int, int, bool) {
	newList := make([]int, 0, len(list)-1)
	for i := 0; i < len(list)-1; i++ {
		d := list[i+1] - list[i]
		newList = append(newList, d)
	}
	for i := 0; i < len(newList)-1; i++ {
		if newList[i] != newList[i+1] {
			return newList, 0, false
		}
	}

	return newList, list[len(list)-1] + newList[0], true
}

// 等比:等比数列是指从第二项起，每一项与它的前一项的比值等于同一个常数的一种数列
func equalRatioSeq(list []int) ([]int, int, bool) {
	newList := make([]int, 0, len(list)-1)
	for i := 0; i < len(list)-1; i++ {
		if list[i] == 0 {
			return []int{}, 0, false
		}
		d := float64(list[i+1]) / float64(list[i])
		if math.Remainder(d, float64(1)) != 0 {
			return []int{}, 0, false
		}
		newList = append(newList, int(d))
	}

	for i := 0; i < len(newList)-1; i++ {
		if newList[i] != newList[i+1] {
			return newList, 0, false
		}
	}
	return newList, list[len(list)-1] * newList[0], true
}

// 开方等差数列:
// 如果一个数列可以对每一项开根号（比如负数，则认为不可以开根号，再比如3，开根号无法得到整数，
// 也认为不可以开根号）并且得到的新的整数数列是等差数列，那么他就是开方等差数列
func squareArithmeticSeq(list []int) ([]int, int, bool) {
	equalDiff := make([]int, 0, len(list))
	for i := 0; i < len(list); i++ {
		d := math.Sqrt(float64(list[i]))
		if math.IsNaN(d) || math.Remainder(d, 1) != 0 {
			return []int{}, 0, false
		}
		equalDiff = append(equalDiff, int(d))
	}
	ls, val, ok := equalDifferenceSeq(equalDiff)
	if !ok {
		return ls, 0, false
	}
	return ls, int(math.Pow(float64(val), 2)), true
}
