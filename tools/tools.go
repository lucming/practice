package tools

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func MaxValue(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if result < nums[i] {
			result = nums[i]
		}
	}

	return result
}

func Abs(val int) int {
	if val < 0 {
		val = -val
	}

	return val
}

func Reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
