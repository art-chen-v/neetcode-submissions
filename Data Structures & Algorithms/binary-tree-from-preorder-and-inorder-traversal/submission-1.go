/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// preorder = [1,2,3,4]
// inorder = [2,1,3,4]

func buildTree(preorder []int, inorder []int) *TreeNode {
	indices := make(map[int]int)
	for i, num := range inorder {
		indices[num] = i
	}

	var dfs func(preorder []int, inorder[]int) *TreeNode
	dfs = func(preorder []int, inorder[]int) *TreeNode {
		if len(preorder) == 0 || len(inorder) == 0 {
			return nil
		}
		root := &TreeNode{Val: preorder[0]}
		if len(preorder) > 1 {
			m := indices[root.Val]
			root.Left = buildTree(preorder[1:m+1], inorder[:m])
			root.Right = buildTree(preorder[m+1:], inorder[m+1:])
		}
		return root
	}

	return dfs(preorder, inorder)
}
