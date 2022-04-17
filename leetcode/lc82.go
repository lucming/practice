package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, last *ListNode
	p := head
	for p != nil {
		if p.Next == nil || (p.Next != nil && p.Val != p.Next.Val) {
			s := &ListNode{
				Val: p.Val,
			}
			if newHead == nil {
				newHead = s
				last = newHead
			} else {
				last.Next = s
				last = last.Next
			}
		} else {
			for p != nil && p.Next != nil && p.Val == p.Next.Val {
				p = p.Next
			}
		}
		p = p.Next
	}

	return newHead
}
