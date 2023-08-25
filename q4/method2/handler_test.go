package method2

import (
	"fmt"
	"testing"
)

func TestHandler(t *testing.T) {
	//res := findMedianSortedArrays([]int{1, 3}, []int{2})
	//fmt.Println(res)
	res2 := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(res2)
}
