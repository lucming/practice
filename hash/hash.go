package hash

import "sort"

//判断两个string是否为字母异位词
func isAnagram(s string, t string) bool {
	m := [26]int{}
	for _, i := range s {
		m[i-rune('a')]++
	}
	for _, i := range t {
		m[i-rune('a')]--
	}

	return m == [26]int{}
}

//两数组交集
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0)
	result := []int{}

	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}

	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			result = append(result, v)
			delete(set, v)
		}
	}

	return result
}

//判断一个数是否为快乐数
//示例：
//输入：19
//输出：true
//解释：
//1^2 + 9^2 = 82
//8^2 + 2^2 = 68
//6^2 + 8^2 = 100
//1^2 + 0^2 + 0^2 = 1
//注意：有可能会出现无限循环，通过map判断是否出现过某个数，来退出循环
func isHappy(n int) bool {
	set := make(map[int]struct{}, 0)
	cur := n
	set[n] = struct{}{}
	for {
		sum := 0
		for cur != 0 {
			sum += (cur % 10) * (cur % 10)
			cur /= 10
		}
		if sum == 1 {
			return true
		}
		cur = sum
		if _, ok := set[cur]; ok {
			return false
		}
		set[cur] = struct{}{}
	}

	return false
}

//两数之和
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//示例:
//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for idx, v := range nums {
		if _, ok := set[target-v]; ok {
			return []int{idx, set[target-v]}
		} else {
			set[v] = idx
		}
	}

	return []int{-1, -1}
}

//四数相加
//给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
//为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。
//例如:
//输入: A = [ 1, 2] B = [-2,-1] C = [-1, 2] D = [ 0, 2] 输出: 2 解释: 两个元组如下:
//(0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
//(1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	count := 0

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}

	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-v3-v4]
		}
	}

	return count
}

//赎金信
//a b两个字符串，判断b中是否有a的所有字符
//canConstruct("a", "b") -> false
//canConstruct("aa", "ab") -> false
//canConstruct("aa", "aab") -> true
func canConstruct(a, b string) bool {
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)
	for _, v := range a {
		map1[v]++
	}
	for _, v := range b {
		map2[v]++
	}

	for k1, m1 := range map1 {
		if _, ok := map2[k1]; ok && m1 <= map2[k1] {
			continue
		}
		return false
	}

	return true
}

//优化
func canConstruct1(a, b string) bool {
	m := make([]int, 26)
	for _, v := range b {
		m[v-'a']++
	}

	for _, v := range a {
		m[v-'a']--
		if m[v-'a'] < 0 {
			return false
		}
	}

	return true
}

//三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//注意： 答案中不可以包含重复的三元组。
//示例：
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]
//双指针法
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			leftTmp, rightTmp := nums[left], nums[right]
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == leftTmp {
					left++
				}
				for left < right && nums[right] == rightTmp {
					right--
				}
			}
		}
	}

	return result
}
