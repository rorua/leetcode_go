package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// iteratively
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

// iteratively2
func reverseList2(head *ListNode) *ListNode {
	var (
		prev *ListNode = nil
		cur            = head
	)

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// iteratively4
func reverseList4(head *ListNode) *ListNode {

	var prev *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

// recursion
func reverseList3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head
	if head.Next != nil {
		newHead = reverseList3(head.Next)
		head.Next.Next = head
	}
	head.Next = nil
	return newHead
}
