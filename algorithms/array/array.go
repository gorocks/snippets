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

// BinaryFirstSearch 二分查找(有重复则返回第一个), 找到返回其位置, 否则返回 -1.
func BinaryFirstSearch(a []int, k int) int {
	n := len(a)
	l, h := 0, n-1
	for l <= h {
		m := l + (h-l)/2
		if k > a[m] {
			l = m + 1
		} else if k < a[m] {
			h = m - 1
		} else {
			if m > 0 && a[m-1] == k {
				h = m - 1 // 继续向左二分查找
			} else {
				return m
			}
		}
	}
	return -1
}

// BinaryLastSearch 二分查找(有重复则返回最后一个), 找到返回其位置, 否则返回 -1
func BinaryLastSearch(a []int, k int) int {
	n := len(a)
	l, h := 0, n-1
	for l <= h {
		m := l + (h-l)/2
		if k > a[m] {
			l = m + 1
		} else if k < a[m] {
			h = m - 1
		} else {
			if m < n-1 && a[m+1] == k {
				l = m + 1 // 继续向右二分查找
			} else {
				return m
			}
		}
	}
	return -1
}
