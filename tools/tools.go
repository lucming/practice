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
