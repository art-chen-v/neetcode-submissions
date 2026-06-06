/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// d - 1 - 2 - 3 - 4
             
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	curr1 := dummy
	curr2 := head
	for n > 0 {
		curr2 = curr2.Next
		n--
	}
	for curr2 != nil {
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
	curr1.Next = curr1.Next.Next
	return dummy.Next
}
