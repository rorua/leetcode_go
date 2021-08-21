package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		remain = 0
		l3     = new(ListNode)
		cl3    = l3
	)

	for l1 != nil || l2 != nil || remain > 0 {
		v1 := 0
		v2 := 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		tmp := v1 + v2
		remain = tmp / 10

		var newNode = new(ListNode)
		newNode.Val = tmp % 10
		l3.Next = newNode
		l3 = newNode
	}

	return cl3.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var l1 = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	var l2 = ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	l3 := addTwoNumbers(&l1, &l2)

	fmt.Printf("%v", l3)

	return
}
