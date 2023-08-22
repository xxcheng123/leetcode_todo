package q56

import (
	"fmt"
	"testing"
)

func TestInnerMerge(t *testing.T) {
	arr := [][]int{
		{1, 3},
		{2, 3},
		{8, 10},
		{15, 18},
		{18, 19},
		{16, 18},
		{15, 20},
		{18, 20},
		{18, 199},
	}
	res := merge(arr)
	fmt.Println(res)
}

func TestInnerMerge2(t *testing.T) {
	arr := [][]int{
		{1, 4}, {
			4, 5,
		},
	}
	res := merge(arr)
	fmt.Println(res)
}
