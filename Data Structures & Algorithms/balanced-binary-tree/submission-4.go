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
		left, leftB := dfs(root.Left)
		right, rightB := dfs(root.Right)
		isBalanced := leftB && rightB && abs(left, right) <= 1

		return 1 + max(left, right), isBalanced
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
