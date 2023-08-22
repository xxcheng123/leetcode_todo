package tools

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(arr ...any) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	var tree = make([]*TreeNode, 0, len(arr))
	rootVal, _ := arr[0].(int)
	root := &TreeNode{
		Val: rootVal,
	}
	tree = append(tree, root)

	for i, v := range arr[1:] {
		if iv, ok := v.(int); ok {
			node := &TreeNode{
				Val: iv,
			}
			tree = append(tree, node)
			if i%2 == 0 {
				tree[i/2].Left = node
			} else {
				tree[i/2].Right = node
			}
		}
	}
	return tree[0]
}
