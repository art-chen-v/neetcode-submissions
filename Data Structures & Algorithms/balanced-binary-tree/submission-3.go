/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) (int, bool)
	dfs = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		l, leftBalanced := dfs(root.Left)
		r, rightBalanced := dfs(root.Right)
		isBalanced := leftBalanced && rightBalanced && abs(l, r) <= 1
		return 1 + max(l, r), isBalanced
	}
	_, ok := dfs(root)
	return ok
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
