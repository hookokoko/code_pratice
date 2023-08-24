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

// 25. K个一组翻转链表
// 1. 一个方法就是下面的这种；
// 2. 还有另一种就是reverse之后返回新的start和end，前提是reverse之前需要断链
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	// 思考下为啥start是循环内的局部变量，而end在循环外
	// 或者说start如果是全局的，会有什么问题? 好像也没什么问题？其实，就是没必要而已
	end := dummy
	start := dummy

	// 为啥终止条件是end.Next != nil
	for end.Next != nil {
		for i := k; i > 0 && end != nil; i-- {
			end = end.Next
		}
		if end == nil {
			break
		}

		start = pre.Next // 问什么这里是for里面的局部变量？

		next := end.Next // 这个很明显，是记录的作用

		end.Next = nil // 逆转前需要断链
		pre.Next = reverseList(start)

		// 逆转后，将尾巴接上
		start.Next = next

		// 重新赋值pre和end，这时start是丢弃的
		pre = start
		end = pre
	}

	return dummy.Next
}

// 148. 排序链表
// 使用归并排序算法，又分自底向上和自顶向下，这两者时间复杂度不同。
// 完成这道题的前提是，合并两个排序链表。
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		// 每一轮循环找到两个待归并的链表
		for cur != nil {
			// 找第一个链表
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 找第二个链表
			head2 := cur.Next
			cur.Next = nil // 需要跟第三个链表断链(如果有)
			cur = head2    // cur在找第一个和第二个链表时是可以共用的中间变量
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 记录下，第三个链表的头(如果cur不是nil)
			var next *ListNode
			if cur != nil {
				next = cur.Next // 这里记录了
				cur.Next = nil  // 这里和第二个链表断链
			}
			// 一切都是为了合并。。。
			prev.Next = mergeSortedList1(head1, head2)

			// 注意是for，找到合并后链表的尾节点，为了下一轮合并
			for prev.Next != nil {
				prev = prev.Next
			}
			// 同时调整cur
			// cur和prev能共用一个吗？不能因为cur和prev是断开的关系
			cur = next
		}
	}
	return dummyHead.Next
}

// 递归合并
func mergeSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeSortedList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeSortedList(l1, l2.Next)
		return l2
	}
}

// 合并非递归
func mergeSortedList1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, l1, l2
	for temp1 != nil && temp2 != nil {
		if temp1.Val > temp2.Val {
			temp.Next = temp2
			temp2 = temp2.Next
		} else {
			temp.Next = temp1
			temp1 = temp1.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else {
		temp.Next = temp2
	}
	return dummy.Next
}

// 23. 合并k个升序链表
// 基础就是合并链表
//  1. 顺序合并，sum+=num这种
//  2. 两两合并，类似排序链表的归并，但又不完全一样，因为这是有多个链表，不太好分。只能是利用递归了。
//     递归的思路是：每次分一半，递归的函数签名类似这种：l1 = merge(list, left, right), l2= merge(list, left, right); 然后按照合并两个链表的方式合并l1、l2
//     但是也可以不用递归，下面写得是个不是递归的方法。
//  3. 利用k容量大小的优先级队列
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	nums := length
	for nums != 1 {
		idx := 0
		for i := 0; i < nums; i += 2 {
			if i+1 == nums {
				lists[idx] = lists[i]
			} else {
				lists[idx] = mergeSortedList1(lists[i], lists[i+1])
			}
			idx++
		}
		nums = idx
	}
	return lists[0]
}
