package q56

//56. 合并区间
// https://leetcode.cn/problems/merge-intervals/description/
/*
思路，先排序，按照x1<x2 然后按照 y1<y2

**/

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 1 {
		return intervals
	}
	N := len(intervals)
	mergeSort(intervals, 0, N-1, 0)
	sameIndex := -1
	for i, v := range intervals {
		if sameIndex == -1 {
			sameIndex = i
		} else if v[0] != intervals[sameIndex][0] {
			mergeSort(intervals, sameIndex, i-1, 1)
			sameIndex = i
		}
	}
	if intervals[sameIndex][0] == intervals[N-1][0] {
		mergeSort(intervals, sameIndex, N-1, 1)
	}
	var arr [][]int = append(make([][]int, 0), intervals[0])
	for _, interval := range intervals[1:] {
		lastInterval := arr[len(arr)-1]
		if lastInterval[0] == interval[0] {
			arr[len(arr)-1] = interval
		} else if lastInterval[1] >= interval[0] {
			if lastInterval[1] < interval[1] {
				arr[len(arr)-1] = []int{lastInterval[0], interval[1]}
			}
		} else {
			arr = append(arr, interval)
		}
	}
	return arr
}

func mergeSort(arr [][]int, left, right, sortIndex int) {
	if left == right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(arr, left, mid, sortIndex)
	mergeSort(arr, mid+1, right, sortIndex)
	innerMerge(arr, left, right, sortIndex)
}

func innerMerge(arr [][]int, left, right, sortIndex int) {
	narr := make([][]int, right-left+1)
	i := 0
	mid := left + (right-left)/2
	p1, p2 := left, mid+1
	for p1 <= mid && p2 <= right {
		if arr[p1][sortIndex] < arr[p2][sortIndex] {
			narr[i] = arr[p1]
			p1++
		} else {
			narr[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 <= mid {
		narr[i] = arr[p1]
		p1++
		i++
	}
	for p2 <= right {
		narr[i] = arr[p2]
		p2++
		i++
	}
	for i, v := range narr {
		arr[i+left] = v
	}
}
