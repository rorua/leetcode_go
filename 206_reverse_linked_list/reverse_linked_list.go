package reverse_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var (
		cur            = head
		next           = cur.Next
		prev *ListNode = nil
	)

	for next != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
