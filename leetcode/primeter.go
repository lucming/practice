package leetcode

type point struct {
	x int
	y int
}

//计算每个格子个周长，有可能是0，1，2，3，4，这个需要根据判断当前格子的上下左右是否存在过同色的格子判断
func calculatePrimeter(arr [][]int, p point) int {
	result := 0
	if arr[p.x][p.y-1] == 0 {
		result++
	}
	if arr[p.x][p.y+1] == 0 {
		result++
	}
	if arr[p.x-1][p.y] == 0 {
		result++
	}
	if arr[p.x+1][p.y] == 0 {
		result++
	}

	return result
}

//求point所在连续区域的周长
func perimeter(arr [][]int, p point) int {
	if p.x == len(arr) || p.x < 0 {
		return 0
	}

	if p.x == len(arr[0]) || p.y < 0 {
		return 0
	}

	if arr[p.x][p.y] == 1 {
		arr[p.x][p.y] = 2
		currentPrimeter := calculatePrimeter(arr, p)
		//从当前点去递归展开去计算挨着的格子的周长
		return currentPrimeter + perimeter(arr, point{
			x: p.x + 1,
			y: p.y,
		}) + perimeter(arr, point{
			x: p.x - 1,
			y: p.y,
		}) + perimeter(arr, point{
			x: p.x,
			y: p.y - 1,
		}) + perimeter(arr, point{
			x: p.x,
			y: p.y + 1,
		})
	}

	return 0
}
