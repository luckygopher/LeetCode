package main

import (
	"testing"
)

func TestCheckList(t *testing.T) {
	data := []struct {
		List []int
		Ok   bool
		Next int
	}{
		{[]int{3, 5, 7, 9, 11}, true, 13},
		{[]int{2, 4, 8, 16, 32}, true, 64},
		{[]int{2, 15, 41, 80}, true, 132},
		{[]int{2, 14, 64, 202, 502, 1062, 2004}, true, 3474},
		{[]int{1, 1, 3, 15, 105, 945}, true, 10395},
		{[]int{1, 3, 7, 9, 11, 0, 1}, false, 0},
		{[]int{1, 4, 9, 16, 17}, false, 0},
	}
	for _, datum := range data {
		if val, ok := CheckList(datum.List); val == datum.Next && ok == datum.Ok {
			t.Logf("val:%d ok:%t\n", val, ok)
		} else {
			t.Errorf("list:%v val:%d ok:%t  errval:%d errok:%t\n", datum.List, datum.Next, datum.Ok, val, ok)
		}
	}
}
