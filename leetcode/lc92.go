package leetcode

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	p := head
	for p != nil {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}

	return pre
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dumHead := &ListNode{Next: head}
	l := dumHead
	for i := 0; i < left-1; i++ {
		l = l.Next
	}

	reverseL := l.Next
	l.Next = nil

	reverseR := reverseL
	for i := 0; i < right-left; i++ {
		reverseR = reverseR.Next
	}

	r := reverseR.Next
	reverseR.Next = nil

	reversed := reverse(reverseL)

	l.Next = reversed
	reverseL.Next = r

	return dumHead.Next
}
