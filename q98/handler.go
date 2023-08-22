package q98

// 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isBST(root).IsBST
}

func isBST(root *TreeNode) *Info {
	if root == nil {
		return nil
	}
	leftInfo := isBST(root.Left)
	rightInfo := isBST(root.Right)
	leftIsBST := true
	rightIsBST := true
	isBst := true
	if leftInfo != nil && !leftInfo.IsBST {
		isBst = false
	}
	if rightInfo != nil && !rightInfo.IsBST {
		isBst = false
	}

	if leftInfo != nil && leftInfo.Max >= root.Val {
		leftIsBST = false
	}
	if rightInfo != nil && rightInfo.Min <= root.Val {
		rightIsBST = false
	}
	if !leftIsBST || !rightIsBST {
		isBst = false
	}
	max, min := root.Val, root.Val
	if leftInfo != nil {
		max = Max(leftInfo.Max, max)
		min = Min(leftInfo.Min, min)
	}
	if rightInfo != nil {
		max = Max(rightInfo.Max, max)
		min = Min(rightInfo.Min, min)
	}
	return &Info{
		Max:   max,
		Min:   min,
		IsBST: isBst,
	}
}

type Info struct {
	Max   int
	Min   int
	IsBST bool
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
