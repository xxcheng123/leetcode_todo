package q327

// 327. 区间和的个数
// https://leetcode.cn/problems/count-of-range-sum/

/**
示例 1：
输入：nums = [-2,5,-1], lower = -2, upper = 2
输出：3
解释：存在三个区间：[0,0]、[2,2] 和 [0,2] ，对应的区间和分别是：-2 、-1 、2 。
示例 2：

输入：nums = [0], lower = 0, upper = 0
输出：1
*/

func countRangeSum(nums []int, lower int, upper int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	//生成前缀累加和
	for i, v := range nums[1:] {
		arr[i+1] = arr[i] + v
	}
	return calCount(arr, 0, len(arr)-1, lower, upper)

}
func calCount(arr []int, left, right, lower, upper int) int {
	if left == right {
		if arr[left] >= lower && arr[left] <= upper {
			return 1
		}
		return 0
	}
	mid := left + (right-left)/2
	return calCount(arr, left, mid, lower, upper) +
		calCount(arr, mid+1, right, lower, upper) +
		merge(arr, left, right, lower, upper)
}
func merge(arr []int, left, right, lower, upper int) int {
	count := 0
	windowL, windowR := left, left
	mid := left + (right-left)/2
	for i := mid + 1; i <= right; i++ {
		innerLower := arr[i] - upper
		innerUpper := arr[i] - lower
		// 1 2 2 2 2
		// 在第一个2就要退出了，所以<
		for windowL <= mid && arr[windowL] < innerLower {
			windowL++
		}
		// 1 2 2 2 2 3
		// 一直到最后应该2才退出 所以<=
		for windowR <= mid && arr[windowR] <= innerUpper {
			windowR++
		}
		count += windowR - windowL
	}
	//正常的merge
	helpArr := make([]int, right-left+1)
	p1, p2 := left, mid+1
	helpIndex := 0
	for p1 <= mid && p2 <= right {
		if arr[p1] < arr[p2] {
			helpArr[helpIndex] = arr[p1]
			p1++
		} else {
			helpArr[helpIndex] = arr[p2]
			p2++
		}
		helpIndex++
	}
	for p1 <= mid {
		helpArr[helpIndex] = arr[p1]
		p1++
		helpIndex++
	}

	for p2 <= right {
		helpArr[helpIndex] = arr[p2]
		p2++
		helpIndex++
	}
	for i, v := range helpArr {
		arr[left+i] = v
	}
	return count
}
