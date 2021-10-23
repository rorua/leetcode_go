package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

//using hashmap
func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	m:= make(map[*ListNode]bool)
	for head != nil {
		if _, found := m[head]; found {
			return true
		}
		head = head.Next
	}
	return false
}
