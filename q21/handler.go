package q21

// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
/*
思路：和归并排序的merge一样的
新建一个链表，头指针，
然后再取两个指针，指向起始位置
开始遍历

*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	pHead := new(ListNode)
	p := pHead
	var p1, p2 *ListNode
	if list1 != nil {
		p1 = list1
	}
	if list2 != nil {
		p2 = list2
	}
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		//指针移到下一个
		p = p.Next
		p.Next = nil
	}
	for p1 != nil {
		p.Next = p1
		p1 = p1.Next
		p = p.Next
		p.Next = nil
	}

	for p2 != nil {
		p.Next = p2
		p2 = p2.Next
		p = p.Next
		p.Next = nil
	}
	return pHead.Next
}
