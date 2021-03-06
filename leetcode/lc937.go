package leetcode

import "sort"

//我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
//（这里，平面上两点之间的距离是欧几里德距离。）
//你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。
//示例 1：
//输入：points = [[1,3],[-2,2]], K = 1
//输出：[[-2,2]]
//解释：
//(1, 3) 和原点之间的距离为 sqrt(10)，
//(-2, 2) 和原点之间的距离为 sqrt(8)，
//由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
//我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
//示例 2：
//输入：points = [[3,3],[5,-1],[-2,4]], K = 2
//输出：[[3,3],[-2,4]]
//（答案 [[-2,4],[3,3]] 也会被接受。）
//解法1：排序
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}

//解法2：快排
func less(point1, point2 []int) bool {
	return point1[0]*point1[0]+point1[1]*point1[1] < point2[0]*point2[0]+point2[1]*point2[1]
}

func kClosest1(points [][]int, k int) [][]int {
	var quickSelect func(int, int)
	quickSelect = func(left, right int) {
		if left < right {
			i, j := left, right
			tmp := points[i]
			for i < j {
				for i < j && less(tmp, points[i]) {
					j--
				}
				points[i] = points[j]

				for i < j && !less(tmp, points[i]) {
					i++
				}
				points[j] = points[i]
			}
			points[i] = tmp
			quickSelect(left, i-1)
			quickSelect(i+1, right)
		}
	}
	quickSelect(0, len(points)-1)
	return points[:k]
}

//解法3：大根堆
