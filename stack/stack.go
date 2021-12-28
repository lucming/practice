package stack

type stack struct {
	queue1 []int
	queue2 []int
}

func (s *stack) Push(val int) {
	s.queue1 = append(s.queue1, val)
}

func (s *stack) Pop() int {
	s.queue2 = append(s.queue2, s.queue1...)
	top := s.queue2[len(s.queue2)-1]
	s.queue2 = s.queue2[:len(s.queue2)-1]
	s.queue1 = s.queue2
	s.queue2 = []int{}

	return top
}
