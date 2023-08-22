package q112

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/description/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return findPathSum(root, targetSum, 0)
}

func findPathSum(node *TreeNode, targetSum, preSum int) bool {
	//叶子节点
	if node.Left == nil && node.Right == nil {
		if preSum+node.Val == targetSum {
			return true
		}
	}

	if node.Left != nil {
		if findPathSum(node.Left, targetSum, preSum+node.Val) {
			return true
		}
	}
	if node.Right != nil {
		if findPathSum(node.Right, targetSum, preSum+node.Val) {
			return true
		}
	}

	return false
}
