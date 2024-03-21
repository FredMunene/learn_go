package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: Node has no children (leaf node)
	if node.Left == nil && node.Right == nil {
		return deleteLeafNode(root, node)
	}

	// Case 2: Node has one child
	if node.Left == nil || node.Right == nil {
		return deleteNodeWithOneChild(root, node)
	}

	// Case 3: Node has two children
	return deleteNodeWithTwoChildren(root, node)
}

func deleteLeafNode(root, node *TreeNode) *TreeNode {
	// If node is the root and it's the only node in the tree
	if node.Parent == nil {
		return nil
	}

	// Update parent pointer of the node's parent to nil
	if node == node.Parent.Left {
		node.Parent.Left = nil
	} else {
		node.Parent.Right = nil
	}

	return root
}

func deleteNodeWithOneChild(root, node *TreeNode) *TreeNode {
	// Identify the non-nil child node
	var child *TreeNode
	if node.Left != nil {
		child = node.Left
	} else {
		child = node.Right
	}

	// If node is the root
	if node.Parent == nil {
		root = child
		child.Parent = nil
		return root
	}

	// Update parent pointer of node's parent to child
	if node == node.Parent.Left {
		node.Parent.Left = child
	} else {
		node.Parent.Right = child
	}
	child.Parent = node.Parent

	return root
}

func deleteNodeWithTwoChildren(root, node *TreeNode) *TreeNode {
	// Find the minimum node in the right subtree
	minNode := findMinNode(node.Right)

	// Copy the data of the minimum node to the node to be deleted
	node.Data = minNode.Data

	// Delete the minimum node (which now has duplicate data) recursively
	root = BTreeDeleteNode(root, minNode)

	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
