package leetcode

func isCycle(data int) bool {
	if data < 0 {
		return false
	}

	tmp := 0
	for data > tmp {
		tmp = tmp*10 + data%10
		data /= 10
	}

	return data == tmp || data == tmp/10
}
