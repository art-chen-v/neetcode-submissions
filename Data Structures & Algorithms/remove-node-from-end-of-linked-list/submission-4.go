/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
    tail := curr
	for n >= 0 {
		tail = tail.Next
		n--
	}
	for tail != nil {
		curr = curr.Next
		tail = tail.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}
