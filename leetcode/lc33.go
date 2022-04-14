package leetcode

//旋转数组查找
func search(nums []int, target int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return -1
	}
	if lenNums == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		}

		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
