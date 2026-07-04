/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
    startNode := dummyNode
	endNode := startNode

	for n >= 0 {
		endNode = endNode.Next
		n--
	}

	for endNode != nil {
		startNode = startNode.Next
		endNode = endNode.Next
	}

	startNode.Next = startNode.Next.Next
	return dummyNode.Next
}
