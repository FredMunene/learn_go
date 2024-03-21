package piscine

// returns number of levels of binary tree ( height of the tree)

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depthleft := BTreeLevelCount(root.Left)
	depthright := BTreeLevelCount(root.Right)

	if depthleft >= depthright {
		return depthleft + 1
	} else {
		return depthright + 1
	}
}
