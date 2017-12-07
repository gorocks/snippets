package tree

// Tree is a binary tree.
type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

// InOrder LDR
func (t *Tree) InOrder() []int {
	a := make([]int, 0)
	if t != nil {
		a = append(a, t.Left.InOrder()...)
		a = append(a, t.Value)
		a = append(a, t.Right.InOrder()...)
	}
	return a
}

func generateTree(a []int, start, end int) *Tree {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	head := &Tree{Value: a[mid]}
	head.Left = generateTree(a, start, mid-1)
	head.Right = generateTree(a, mid+1, end)
	return head
}

// NewInOrderTree 通过一个有序数组构造一个中序二分查找树.
// 中序 LDR 即左根右顺序.
func NewInOrderTree(a []int) *Tree {
	length := len(a)
	if length == 0 {
		return nil
	}
	return generateTree(a, 0, length-1)
}
