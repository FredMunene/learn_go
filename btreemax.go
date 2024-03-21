package piscine

// returns the node with the maximum value in the tree given by root

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Right == nil { // if right child at current node is nil
		return root
	}
	return BTreeMax(root.Right)
}
