package leetcode

type node struct {
	left  *node
	right *node
	value int
}

func getCommonAncestor(root, p, q *node) *node {
	if root == p || root == q || root == nil {
		return root
	}

	left := getCommonAncestor(root.left, p, q)
	right := getCommonAncestor(root.right, p, q)

	if left == nil && right == nil {
		return nil
	} else if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	}

	return root
}
