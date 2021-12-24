package search

//给你一个数组 nums 和一个值 val，原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
func removeElement(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}

	return slow
}

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
//示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]
//思路：因为数组有序，所以平方的结果最大值肯定在两边，所以使用快慢指针，i,j分别指向数组的开始和结尾，谁的平方值小就将该下标的值的平方写到结果数组中
func sortedSquares(arr []int) []int {
	ret := make([]int, len(arr))
	idx := len(arr) - 1
	i, j := 0, len(arr)-1
	for i <= j {
		if arr[i]*arr[i] > arr[j]*arr[j] {
			ret[idx] = arr[i] * arr[i]
			i++
		} else {
			ret[idx] = arr[j] * arr[j]
			j--
		}
		idx--
	}

	return ret
}

//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
//示例：
//输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//思路： 滑动窗口：就是不断的调节子序列的起始位置和终止位置，从而得出我们要想的结果
func minSubArrayLen(nums []int, s int) int {
	start, end := 0, 0
	minLen := 0
	for end < len(nums) {
		if sum(nums[start:end+1]) < s {
			end++
		} else {
			if minLen == 0 {
				minLen = end - start + 1
			} else {
				if minLen > end-start+1 {
					minLen = end - start + 1
				}
			}
			start++
		}

	}

	return minLen
}

func sum(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}

	return result
}
