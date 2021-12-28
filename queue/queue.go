package stack

//两个栈实现队列
//思路：入s2 出s1
type queue struct {
	stack1 []int
	stack2 []int
}

//入队
func (q *queue) Push(val int) {
	for len(q.stack1) > 0 {
		q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
		q.stack1 = q.stack1[:len(q.stack1)-1]
	}
	q.stack2 = append(q.stack2, val)
}

//出队
func (q *queue) Pop() int {
	for len(q.stack2) > 0 {
		q.stack1 = append(q.stack1, q.stack2[len(q.stack2)-1])
		q.stack2 = q.stack2[:len(q.stack2)-1]
	}
	if len(q.stack1) == 0 {
		return 0
	}
	out := q.stack1[len(q.stack1)-1]
	q.stack1 = q.stack1[:len(q.stack1)-1]

	return out
}

//获取队头元素
//要不在stack1尾，要不在stack2头
func (q *queue) Peek() int {
	if len(q.stack2) > 0 {
		return q.stack2[0]
	}

	if len(q.stack1) > 0 {
		return q.stack1[len(q.stack1)-1]
	}

	return 0
}
