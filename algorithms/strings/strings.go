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

// NumSum 字符串中数字子串的求和
// 给定的一个字符串 s, 求其中全部数字串所代表的数字之和.
// * 忽略小数点字符, 例如 "A1.3", 其中包含两个数 1 和 3.
// * 如果紧贴数字子串的左侧出现字符 "-", 当连续出现的数量为奇数时, 则数字视为负,
// 连续出现的数量为偶数时, 则数字视为正. 例如 "A-1BC--12", 其中包含数字为 -1 和 12.
func NumSum(s string) int {
	// if len(s) == 0 {
	// 	return 0 A1CD2E33 A-1B--2C--D6E
	// }
	var sum, temp int // sum: 累加和 temp: 中间临时数
	pos := true       // 正或负
	for _, v := range s {
		if v == '-' {
			pos = !pos
		} else {
			if vv := v - '0'; vv >= 0 && vv <= 9 {
				temp = 10*temp + int(vv)
			} else {
				if !pos {
					temp = -temp
					pos = true
				}
				sum += temp
				temp = 0
			}
		}
	}
	return sum + temp
}
