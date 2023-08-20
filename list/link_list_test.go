package list

import (
	"fmt"
	"testing"
)

func Test_ReverseList(t *testing.T) {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}

	res := reverseList(&l1)
	fmt.Println(printList(res))
}

func Test_AddTwoNumbers(t *testing.T) {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}

	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 8}
	l2.Next.Next = &ListNode{Val: 9}
	//l2.Next.Next.Next = &ListNode{Val: 8}

	res := addTwoNumbers(&l1, &l2)
	fmt.Println(printList(res))
}

func Test_SwapPairs(t *testing.T) {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}

	res := swapPairs(&l1)
	fmt.Println(printList(res))
}

func printList(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}
