package dynamic

//斐波那契数列
//f(n)=f(n-1)+f(n-2)
func fib(n int) int {
	if n < 2 {
		return n
	}

	morepre, pre, current := 0, 1, 0
	for i := 2; i <= n; i++ {
		current = pre + morepre
		morepre = pre
		pre = current
	}

	return current
}
