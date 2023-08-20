package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 坑：想一想这么赋值会有什么问题？最后一位是1，1会指向2，2又指向1，死循环了。
	//pre := head
	//cur := head.Next

	// 正确的赋值方法
	var pre *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 求两个的链表元素的加和, 其中链表元素是就是数字中的1位，即不会超过9，然后链表是数字的逆序
// 如数字123，链表就是3→2→1, 这样求加和就比较方便，相当于从个位开始加了。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	var carry int
	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10

		// 算法的难点在于 pre和cur的处理
		// 尝试过只定义cur，这样新链表就找不到了，因为cur一直会往后遍历的
		// 相当于通过pre标记新链表的head
		// 只为为啥返回的是pre.Next，还是和实现有关
		newNode := &ListNode{Val: sum}
		cur.Next = newNode
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		next := new(ListNode)
		(*next).Val = 1
		cur.Next = next
	}

	return pre.Next
}

/*
两两交换链表中的元素
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy

	// e.g. dummy → 1 → 2 → 3 → 4
	for cur.Next != nil && cur.Next.Next != nil {
		// 需要记录的是1、3节点
		tmp1 := cur.Next
		tmp2 := cur.Next.Next.Next
		// 先让dummy指向2
		cur.Next = tmp1.Next
		// 再让2指向1
		cur.Next.Next = tmp1
		// 最后让1指向3
		tmp1.Next = tmp2
		// 修改cur的指针，因为是两两交换所以前进2步
		cur = cur.Next.Next
	}
	return dummy.Next
}
