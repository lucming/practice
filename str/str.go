package str

//字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
//示例 1：
//输入: s = "abcdefg", k = 2
//输出: "cdefgab"
//示例 2：
//输入: s = "lrloseumgh", k = 6
//输出: "umghlrlose"
//
//思路： 部分旋转再整体旋转，首先旋转0-k,再旋转k+1-结尾，再整体旋转
//eg:
//s=abcdef,k=2
//bacdef
//bafedc
//cdefab
func reverseLeftWords(s string, k int) string {
	byteStr := []byte(s)
	reverse(byteStr, 0, k-1)
	reverse(byteStr, k, len(byteStr)-1)
	reverse(byteStr, 0, len(byteStr)-1)

	return string(byteStr)
}

func reverse(s []byte, left int, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
