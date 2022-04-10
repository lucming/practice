package leetcode

import "strconv"

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
func searchMatrix(matrix [][]int, target int) bool {
	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}

	return false
}

func MaxNumByDel3(num int) int {
	strNum := strconv.Itoa(num)
	delIdx := 0
	last3Idx := 0
	for idx, s := range strNum {
		if s != '3' {
			continue
		}
		last3Idx = idx
		if idx == len(strNum)-1 {
			break
		}
		if strNum[0] == '-' {
			if strNum[idx] > strNum[idx+1] {
				delIdx = idx
				break
			}
		} else {
			if strNum[idx] < strNum[idx+1] {
				delIdx = idx
				break
			}
		}
	}

	if delIdx == 0 {
		delIdx = last3Idx
	}
	strNum = strNum[:delIdx] + strNum[delIdx+1:]
	result, err := strconv.Atoi(strNum)
	if err != nil {
		return 0
	}

	return result
}
