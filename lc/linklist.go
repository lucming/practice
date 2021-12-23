package lc

import "fmt"

type node struct {
	data int
	next *node
}

func (n *node) headInsert(d int) {
	s := &node{
		data: d,
	}
	if n == nil {
		n = s
	} else {
		s.next = n.next
		n.next = s
	}
}

func (n *node) tailInsert(d int) {
	p := n
	for ; p.next != nil; p = p.next {
		p = p.next
	}
	s := &node{data: d}
	s.next = p.next
	p.next = s
}

//头插法实现链表逆置
func reverse(head *node) {
	if head == nil || head.next == nil || head.next.next == nil {
		return
	}

	p := head.next
	head.next = nil
	var q *node
	for p != nil {
		q = p.next
		p.next = head.next
		head.next = p
		p = q
	}
}

//结点不动，第一个结点的数据域与最后一个结点的数据域的值交换，
//第二个结点的数据域与倒数第二个结点的数据域交换，以此类推
//时间复杂度O(n^2)
func reverse1(head *node) {
	p := head.next
	var q *node
	for q = head; q.next != nil; q = q.next {
		for i := 0; i < getLength(head)/2; i++ {
			p.data, q.data = q.data, p.data
			p = p.next
			q = getPri(head, q)
		}
	}
}

//改变指针的指向，后->前，头指针指向最后一个结点 //时间复杂的O(n)
func reverse2(head *node) {
	if head == nil || head.next == nil || head.next.next == nil {
		return
	}

	p := head.next
	q := p.next
	p.next = nil
	var s *node

	for q != nil {
		s = q.next
		q.next = p
		p = q
		q = s
	}
	head.next = p
}

func getLength(head *node) int {
	count := 0
	for head != nil {
		head = head.next
		count++
	}

	return count
}

func getPri(head *node, find *node) *node {
	if head == nil {
		return head
	}

	for head.next != nil && head.next != find {
		head = head.next
	}

	return head
}

func (n *node) show() {
	for p := n; p != nil; p = p.next {
		fmt.Println(p.data)
	}
}

func cycle(head *node) bool {
	if head == nil {
		return false
	}

	slow, quick := head, head
	for quick != nil && quick.next != nil {
		slow = slow.next
		quick = quick.next.next
		if slow == quick {
			return true
		}
	}

	return false
}

//给定单链表(head)，如果有环的话请返回从头结点进入环的第一个节点。
//先判断是否有环。
//如果有环，那么p1p2重合点p必然在环中。从p点断开环，方法为：p1=p, p2=p->next, p->next=NULL。此时，原单链表可以看作两条单链表，一条从head开始，另一条从p2开始，我们找到它们的第一个交点即为所求。
func getLoppNode(head *node) *node {
	if head == nil {
		return nil
	}

	slow, quick := head, head
	for quick != nil && quick.next != nil {
		slow = slow.next
		quick = quick.next.next
		if slow == quick {
			slow = head
			for {
				quick = quick.next
				slow = slow.next
				if quick == slow {
					break
				}
			}
			return slow
		}
	}

	return nil
}

//给定两个单链表(head1, head2)，检测两个链表是否有交点，如果有返回第一个交点。
//方法一：
//如果head1==head2，那么显然相交，直接返回head1。
//否则，分别从head1,head2开始遍历两个链表获得其长度len1与len2。假设len1>=len2，那么指针p1由head1开始向后 移动len1-len2步。指针p2=head2，下面p1、p2每次向后前进一步并比较p1p2是否相等，如果相等即返回该结点，否则说明两个链表没有 交点。
//方法二：
//将其中一条单链表首尾相连，转化成求链表入环的第一个结点
func getPoint(head1, head2 *node) *node {
	p := head1
	for p.next != nil {
		p = p.next
	}
	p.next = head1

	return getLoppNode(head2)
}

//只给定单链表中某个结点p(并非最后一个结点，即p->next!=NULL)指针，删除该结点。
//首先是放p中数据,然后将p->next的数据copy入p中，接下来删除p->next即可。
func deletePoint(head *node, find *node) {
	if head == nil {
		return
	}

	p := head
	for p.next != nil {
		if p.next == find {
			p.next = find.next
		}
		p = p.next
	}
}

//只给定单链表中某个结点p(非空结点)，在p前面插入一个结点。
//办法与前者类似，首先分配一个结点q，将q插入在p后，接下来将p中的数据copy入q中，然后再将要插入的数据记录在p中
//只给定单链表中某个结点p（非空）,在p前插一个结点
func insertPri(head, p, s *node) {
	ptmp := head
	for ptmp.next != nil {
		if ptmp.next == p {
			s.next = ptmp.next
			ptmp.next = s
			break
		}
		ptmp = ptmp.next
	}
}

//给定单链表头结点，删除链表中倒数第k个节点
//思路:快慢指针，快指针先走k步，然后快慢指针一起走，快指针走的末尾，慢指针就在第n-1步，然后慢指针跳一个节点
func deleteK(head *node, k int) {
	if getLength(head) < k {
		return
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.next
	}

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}
