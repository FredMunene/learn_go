package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	// recursive search in left
	leftResult := BTreeSearchItem(root.Left, elem)
	// recursive search in right
	rightResult := BTreeSearchItem(root.Right, elem)
	if rightResult != nil {
		return rightResult
	}
	if root.Data == elem {
		return root
	}
	if leftResult != nil {
		return leftResult
	}
	return nil
}
