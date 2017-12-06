package array

// MergeTwoArray 合并两个有序数组(不管有没有重复元素).
func MergeTwoArray(a, b []int) []int {
	an, bn := len(a), len(b)
	c := make([]int, an+bn)
	var i, j, k int
	for i < an && j < bn {
		if a[i] < b[j] {
			c[k] = a[i]
			k++
			i++
		} else {
			c[k] = b[j]
			k++
			j++
		}
	}
	for i < an {
		c[k] = a[i]
		k++
		i++
	}
	for j < bn {
		c[k] = b[j]
		k++
		j++
	}
	return c
}
