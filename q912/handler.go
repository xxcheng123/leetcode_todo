package q912

// 912. 排序数组
// https://leetcode.cn/problems/sort-an-array/description/

func sortArray(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, left, right int) {
	if left == right {
		return
	}
	mid := left + (right-left)/2
	sort(nums, left, mid)
	sort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	help := make([]int, right-left+1)
	p1, p2 := left, mid+1
	helpIndex := 0
	for p1 <= mid && p2 <= right {
		if nums[p1] < nums[p2] {
			help[helpIndex] = nums[p1]
			p1++
		} else {
			help[helpIndex] = nums[p2]
			p2++
		}
		helpIndex++
	}
	for p1 <= mid {
		help[helpIndex] = nums[p1]
		p1++
		helpIndex++
	}
	for p2 <= right {
		help[helpIndex] = nums[p2]
		p2++
		helpIndex++
	}
	for i, v := range help {
		nums[i+left] = v
	}
}
