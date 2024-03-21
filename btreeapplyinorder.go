package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// func applies a given function f,in order, to each element in the tree

func BTreeApplyInorder(root *TreeNode, F func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, F)
	F(root.Data)
	BTreeApplyInorder(root.Right, F)
}
