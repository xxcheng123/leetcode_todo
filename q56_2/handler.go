package q56_2

// merge 思路二，小根堆
func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) <= 1 {
		return intervals
	}
	//先排序
	mergeSort(intervals, 0, len(intervals)-1)
	res := make([][]int, 0)
	h := &heap{}
	lastLeft := intervals[0][0]
	h.insert(intervals[0][1])
	for _, interval := range intervals[1:] {
		a, b := interval[0], interval[1]
		peek := h.peek()
		for h.heapSize != 0 && h.peek() < a {
			peek = h.pop()
		}
		if h.heapSize == 0 {
			res = append(res, []int{lastLeft, peek})
			lastLeft = a
		}
		h.insert(b)
	}
	if h.heapSize != 0 {
		peek := h.peek()
		for h.heapSize != 0 {
			peek = h.pop()
		}
		res = append(res, []int{lastLeft, peek})
	}
	return res
}

type heap struct {
	arr      []int
	heapSize int
}

func (h *heap) insert(val int) {
	if h.heapSize < len(h.arr) {
		h.arr[h.heapSize] = val
	} else {
		h.arr = append(h.arr, val)
	}
	index := h.heapSize
	for h.arr[index] < h.arr[(index-1)/2] {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}

	h.heapSize++

}
func (h *heap) heapify(index int) {
	left := index*2 + 1
	for left < h.heapSize {
		min := left
		if left+1 < h.heapSize && h.arr[left+1] < h.arr[min] {
			min = left + 1
		}
		if h.arr[index] > h.arr[min] {
			h.swap(index, min)
			index = min
			left = index*2 + 1
		} else {
			break
		}
	}
}
func (h *heap) swap(a, b int) {
	temp := h.arr[a]
	h.arr[a] = h.arr[b]
	h.arr[b] = temp
}
func (h *heap) pop() int {
	h.swap(0, h.heapSize-1)
	h.heapSize--
	h.heapify(0)
	return h.arr[h.heapSize]
}
func (h *heap) peek() int {
	return h.arr[0]
}

// 需要一个额外的辅助空间来存放
func mergeSort(intervals [][]int, left, right int) {
	if left == right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(intervals, left, mid)
	mergeSort(intervals, mid+1, right)
	mergeSortWork(intervals, left, right)
}

func mergeSortWork(intervals [][]int, left, right int) {
	mid := left + (right-left)/2
	helpArr := make([][]int, right-left+1)
	p1, p2 := left, mid+1
	i := 0
	for p1 <= mid && p2 <= right {
		if intervals[p1][0] > intervals[p2][0] {
			helpArr[i] = intervals[p2]
			p2++
		} else {
			helpArr[i] = intervals[p1]
			p1++
		}
		i++
	}
	for p1 <= mid {
		helpArr[i] = intervals[p1]
		p1++
		i++
	}
	for p2 <= right {
		helpArr[i] = intervals[p2]
		p2++
		i++
	}
	for i, v := range helpArr {
		intervals[left+i] = v
	}
}
