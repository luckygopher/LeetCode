// @Description:
// @Author: Arvin
// @Date: 2021/4/8 10:36 下午
package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	// 设置矩形边界
	right, bottom := len(grid[0])-1, len(grid)-1
	// 设置队列
	list := make([][3]int, 0)
	list = append(list, [3]int{0, 0, grid[0][0]})

	for len(list) > 0 {
		seed := list[0]
		if seed[0] == right && seed[1] == bottom {
			break
		}
		list = list[1:]
		if seed[0] < right {
			x := seed[0] + 1
			y := seed[1]
			sum := grid[y][x] + seed[2]
			list = append(list, [3]int{x, y, sum})
		}
		if seed[1] < bottom {
			x := seed[0]
			y := seed[1] + 1
			sum := grid[y][x] + seed[2]
			list = append(list, [3]int{x, y, sum})
		}
	}
	fmt.Println(list)
	var result int
	for k, v := range list {
		if k == 0 {
			result = v[2]
		} else if v[2] < result {
			result = v[2]
		}
	}
	return result
}
