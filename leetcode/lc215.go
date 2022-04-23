package leetcode

//第k大的数字
//思路：快排逆序，如果当前的下标是k-1,那么他就是第k大的数字
func findKthLargest(nums []int, k int) int {
	quickSort(nums, k, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort(nums []int, k, start, end int) {
	if start < end {
		i, j := start, end
		tmp := nums[i]
		for i < j {
			for i < j && tmp > nums[j] {
				j--
			}
			nums[i] = nums[j]
			for i < j && tmp <= nums[i] {
				i++
			}
			nums[j] = nums[i]
		}
		nums[i] = tmp
		if i == k-1 {
			return
		} else if i < k-1 {
			quickSort(nums, k, i+1, end)
		} else {
			quickSort(nums, k, start, i-1)
		}
	}
}
