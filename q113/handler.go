package q113

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/description/
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result = make([][]int, 0)
	var matchPaths = make([]int, 0)
	findPathSum(root, targetSum, 0, &matchPaths, &result)
	return result
}

func findPathSum(node *TreeNode, targetSum int, preSum int, matchPaths *[]int, result *[][]int) {
	if node == nil {
		return
	}

	//叶子节点
	if node.Left == nil && node.Right == nil {
		if preSum+node.Val == targetSum {
			*matchPaths = append(*matchPaths, node.Val)
			copiedMathPaths := make([]int, len(*matchPaths))
			copy(copiedMathPaths, *matchPaths)
			*result = append(*result, copiedMathPaths)
			*matchPaths = (*matchPaths)[:len(*matchPaths)-1]
		}
		return
	}
	*matchPaths = append(*matchPaths, node.Val)
	if node.Left != nil {
		findPathSum(node.Left, targetSum, preSum+node.Val, matchPaths, result)
	}

	if node.Right != nil {
		findPathSum(node.Right, targetSum, preSum+node.Val, matchPaths, result)
	}
	*matchPaths = (*matchPaths)[:len(*matchPaths)-1]
}
