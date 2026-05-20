/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var inorder func(root *TreeNode)
	count := 0
	res := 0
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		count++
		if count == k {
			res = root.Val
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}
