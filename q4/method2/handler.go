package method2

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	l := l1 + l2
	p1, p2, index, lastV, currentV := 0, 0, 0, 0, 0
	for index <= l/2 && p1 < l1 && p2 < l2 {
		lastV = currentV
		if nums1[p1] < nums2[p2] {
			currentV = nums1[p1]
			p1++
		} else {
			currentV = nums2[p2]
			p2++
		}
		index++
	}
	for index <= l/2 && p1 < l1 {
		lastV = currentV
		currentV = nums1[p1]
		p1++
		index++
	}
	for index <= l/2 && p2 < l2 {
		lastV = currentV
		currentV = nums2[p2]
		p2++
		index++
	}
	if l%2 == 1 {
		return float64(currentV)
	} else {
		return float64(lastV) + (float64(currentV)-float64(lastV))/2.0
	}
}
