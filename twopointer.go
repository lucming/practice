package main

//数组去重，返回去重后的长度
//eg：[1,1,2]=>2  [0,1,1,2,3,3,3,3,4]=>5
//思路：快慢指针,
//快指针一直往前遍历，慢指针总是指向未重复数组的最后一个下标
//如果快满指针指向的数字不同，那么下一个不重复的数字出现，把它加到不重复数组中，慢指针往后走一个，把快指针指向的数字追加过来
func removeDuplicates(nums []int) int {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}
