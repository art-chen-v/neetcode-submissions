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
		left, leftBalanced := dfs(root.Left)
		right, rightBalanced := dfs(root.Right)
		isBalanced := leftBalanced && rightBalanced && abs(left, right) <= 1
		return max(left, right) + 1, isBalanced
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