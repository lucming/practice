package stack

import "sort"

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

//删除字符串中所有相邻的重复项
//示例：
//输入："abbaca"
//输出："ca"
//解释：bb重复了,消掉,然后aa就相邻了再消掉
func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

//有效的括号
//eg:"[{()}]{}[]" => true
func isValid(s string) bool {
	if s == "" {
		return false
	}
	stack := make([]byte, 0)
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}

	return true
}

//前k个高频词
//两个方法方法，
//1.维护k个节点de小根堆
//2. map统计每个字符出现出现的次数，按照出现次数排序，返回前k个节点
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	allData := make([]int, 0)
	for k, _ := range m {
		allData = append(allData, k)
	}

	sort.Slice(allData, func(a, b int) bool {
		return m[allData[a]] > m[allData[b]]
	})

	return allData[:k]
}

//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//例如，给定一个列表
//temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，
//你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
func dailyTemperatures(nums []int) []int {
	result := make([]int, len(nums))
	stack := []int{}

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = i - top
		}
		stack = append(stack, i)
	}

	return result
}
