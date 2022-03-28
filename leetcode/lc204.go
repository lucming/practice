package leetcode

//给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
//思路：先初始化素数的数组全为true,然后标记出非素数的位置
//标记的思路：如果一个数是素数，那个这个数的倍数肯定就不是素数
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
