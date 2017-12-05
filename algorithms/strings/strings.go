package strings

// IsDeformation 判断两个字符串是否互为变形词
// 给定的两个字符串 s1 和 s2, 如果 s1 和 s2 中出现的字符种类一样且每种字符出现的次数也一样, 那么 s1 和 s2 互为变形词.
// 时间复杂度 O(N), 空间复杂度 O(M)
func IsDeformation(s1, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 == 0 || n2 == 0 || n1 != n2 {
		return false
	}
	m := make(map[rune]int, 0)
	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	for _, v := range s2 {
		if _, ok := m[v]; !ok || m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
