package lc

//问题：长度为n的数组中s所有数字都在0～n-1范围内，数组中某些数字重复了，但不知道有几个数字重复了，也不知道每个数字重复了几次，请找出任意一个重复数字
//eg: input: [2,3,1,0,2,5,3]
// output:2或者3

//思路：重排数组，遍历数组，遍历到i时，判断arr[i]==i,不等则进行交换，入上面的例子，第一次遍历到2,2号下标的数字不是2，交换2号下标数字与i下标的值，模拟该过程
//[1,3,2,0,2,5,3]
//[3,1,2,0,2,5,3]
//[0,1,2,3,2,5,3]
//...
func duplicate(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		for arr[i] != i {
			if arr[i] == arr[arr[i]] {
				return arr[i]
			}
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		}
	}

	return -1
}

//不修改原数组找出数组中重复的数字
//长度为n+1的数组，所有数字都在1-n之间，所以肯定有重复的数字，找出任意一个重复的数字，但是不能修改数组
//思路：将从1-n的中间数字m，将数组分成1-m,m+1,n,如果1-m的数字出现次数超过m,那么这个区间一定包含重复数字，否在在另一个区间，类似二分查找
func getBuplicate(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		c := count(arr, start, mid)
		if start == end {
			if c > 1 {
				return start
			} else {
				break
			}
		}

		if c > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return -1
}

func count(arr []int, left, right int) int {
	count := 0
	for _, val := range arr {
		if left <= val && val <= right {
			count++
		}
	}

	return count
}
