package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	p := head
	if p == nil || p.Next == nil {
		return head
	}

	var pre *ListNode
	p1, p2, p := head, head.Next, head.Next.Next
	newHead := p2
	for p2 != nil {
		p2.Next = p1
		if pre != nil {
			pre.Next = p2
		} else {
			p1.Next = pre
		}
		pre = p1
		p1 = p
		if p1 == nil {
			pre.Next = p1
			return newHead
		}
		p2 = p1.Next
		if p2 == nil {
			pre.Next = p1
			return newHead
		}
		p = p2.Next
	}

	return newHead
}
