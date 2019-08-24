package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2

	root := &ListNode{}
	node := root
	tenFlag := 0

	for (true) {
		var v1, v2 = 0, 0
		if node1 != nil {
			v1 = node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			v2 = node2.Val
			node2 = node2.Next
		}

		node.Val = v1 + v2 + tenFlag
		tenFlag = 0

		if node.Val >= 10 {
			node.Val = node.Val - 10
			tenFlag = 1
		}

		if node1 != nil || node2 != nil || tenFlag != 0 {
			node.Next = &ListNode{}
			node = node.Next
		} else {
			break
		}
	}

	return root
}
