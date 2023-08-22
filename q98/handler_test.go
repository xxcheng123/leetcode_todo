package q98

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {
	root := &TreeNode{
		Val: 32,
		Left: &TreeNode{
			Val: 26,
			Left: &TreeNode{
				Val:  19,
				Left: nil,
				Right: &TreeNode{
					Val: 27,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  47,
			Left: nil,
			Right: &TreeNode{
				Val: 56,
			},
		},
	}
	res := isValidBST(root)
	fmt.Println(res)
}
