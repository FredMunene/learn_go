package piscine

// returns true if only tree given by root folow binary search tree prporties

func BTreeIsBinary(root *TreeNode) bool {
	// left < right
	return BtreeProperty(root, nil, nil)
}

func BtreeProperty(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}
	nodeValue := node.Data
	if (min != nil && nodeValue <= *min) || (max != nil && nodeValue >= *max) {
		return false
	}
	return BtreeProperty(node.Left, min, &node.Data) && BtreeProperty(node.Right, &nodeValue, max)
}
