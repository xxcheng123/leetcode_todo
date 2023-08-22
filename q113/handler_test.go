package q113

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {
	root := BuildTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1)
	res := pathSum(root, 22)
	fmt.Println(res)
}

func TestBuildTree(t *testing.T) {
	res := BuildTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1)
	preorderTraversal(res)

}

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}
