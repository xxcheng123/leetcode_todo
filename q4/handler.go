package q4

// 4. 寻找两个正序数组的中位数
// https://leetcode.cn/problems/median-of-two-sorted-arrays/description/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	l := len(nums1)
	sort(nums1, 0, l-1)
	if l%2 == 1 {
		return float64(nums1[l/2])
	} else {
		a := float64(nums1[l/2])
		b := float64(nums1[l/2-1])
		return a + (b-a)/2
	}
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
