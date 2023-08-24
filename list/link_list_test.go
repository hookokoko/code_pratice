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

func Test_ReverseKGroup(t *testing.T) {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	l1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	res := reverseKGroup(&l1, 3)
	fmt.Println(printList(res))
}

func Test_mergeSortedList(t *testing.T) {
	l1 := generateList([]int{1, 2, 4})
	l2 := generateList([]int{1, 3, 4, 5})
	l := mergeSortedList1(l1, l2)
	fmt.Println(printList(l))
}

func Test_sortList(t *testing.T) {
	l1 := generateList([]int{3, 3, 8, 5, 4, 1, 9})
	l := sortList(l1)
	fmt.Println(printList(l))
}

func Test_mergeKLists(t *testing.T) {
	//[1,4,5],[1,3,4],[2,6]
	l1 := generateList([]int{1, 4, 5})
	l2 := generateList([]int{1, 3, 4})
	l3 := generateList([]int{2, 6})

	l := mergeKLists([]*ListNode{l1, l2, l3})
	fmt.Println(printList(l))
}

func generateList(arr []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for i := 0; i < len(arr); i++ {
		newNode := &ListNode{Val: arr[i]}
		cur.Next = newNode
		cur = cur.Next
	}
	return dummy.Next
}

func printList(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}
