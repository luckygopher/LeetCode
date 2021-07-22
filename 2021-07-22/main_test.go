package main

import (
	"fmt"
	"testing"
)

func TestCheckList(t *testing.T) {
	list := []int{1, 4, 9, 16, 25}
	val, ok := checkList(list)
	fmt.Println(val, ok)
}
