/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	res := 0
    var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			res = root.Val
			return
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}
