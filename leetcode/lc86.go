package leetcode

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	pSmall := small
	large := &ListNode{}
	pLarge := large
	for head != nil {
		if head.Val < x {
			pSmall.Next = head
			pSmall = pSmall.Next
		} else {
			pLarge.Next = head
			pLarge = pLarge.Next
		}
		head = head.Next
	}
	pLarge.Next = nil
	pSmall.Next = large.Next

	return small.Next
}
