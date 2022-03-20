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

//给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
//请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
//示例 1:
//输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出: [-1,3,-1]
//解释:
//对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
//对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
//对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
//示例 2:
//输入: nums1 = [2,4], nums2 = [1,2,3,4].
//输出: [3,-1]
//解释:
//对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
//对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出-1
//思路：先把num1的value和idx做映射，然后遍历num2，
//通过前面的单调栈规律，如果nums2的value中也出现在nums1中
//如果nums2中的当前值>栈顶元素，说明找到了，就把当前值写在nums2在nums1中的下标,
func nextGreaterElement(nums1, nums2 []int) []int {
	m := make(map[int]int)
	for k, v := range nums1 {
		m[v] = k
	}

	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
	}
	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			if _, ok := m[nums2[top]]; ok {
				result[m[nums2[top]]] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}

//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
//示例 1:
//输入: [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；数字 2 找不到下一个更大的数；第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//思路：nums数组遍历两遍，然后有单调栈存下一个最大值
func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = -1
	}

	stack := make([]int, 0)
	for i := 0; i < 2*len(nums); i++ {
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = nums[i%len(nums)]
		}
		stack = append(stack, i%len(nums))
	}

	return result
}
