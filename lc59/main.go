// @Description: 螺旋矩阵
// @Author: Arvin
// @date: 2021/3/22 10:46 上午
package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	// 右下左上跑，跑完就缩圈，直到num越界即停止
	// 创建矩阵,初始值均为0，由于我们是正整数，因此只要访问到不为0则表示已经访问过，应该转向
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// 定以边界
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1

	for num <= n*n {
		for r := left; r <= right; r++ {
			res[top][r] = num
			num++
		}
		top++

		for b := top; b <= bottom; b++ {
			res[b][right] = num
			num++
		}
		right--

		for l := right; l >= left; l-- {
			res[bottom][l] = num
			num++
		}
		bottom--

		for t := bottom; t >= top; t-- {
			res[t][left] = num
			num++
		}
		left++
	}
	return res
}