/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	l := dummyNode
	i := 0
	r := l.Next
	for i < n  && r != nil {
		r = r.Next
		i++
	}

	for r != nil {
		l = l.Next
		r = r.Next
	} 

	l.Next = l.Next.Next

	return dummyNode.Next
}
