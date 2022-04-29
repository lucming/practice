package leetcode

//旋转数组寻找最小数字
//思路：二分查找（要求的数字，要求的数字必定在非有序的数组中，找中间位置，一直卡无序子数组，缩小范围）
func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[l] {
			l = mid + 1
		} else {
			l++
		}
	}

	return nums[l]
}
