package q56_2

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

func TestBuildHeap(t *testing.T) {
	h := &heap{}
	h.insert(9)
	h.insert(2)
	h.insert(8)
	h.insert(5)
	h.insert(4)
	h.insert(3)
	h.insert(7)
	fmt.Println(h)
	h.arr[2] = 99
	h.heapify(2)
	fmt.Println(h)
}
