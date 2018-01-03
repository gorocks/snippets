package defers

func doubleInt(i int) int {
	var r int
	defer func() {
		if r < 1 || r >= 100 {
			r = i
		}
	}()
	r = i * 2
	return r
}

func doubleInt0(i int) (r int) {
	defer func() {
		if r < 1 || r >= 100 {
			r = i
		}
	}()
	return i * 2 // return 此处非原子, 执行步骤: i * 2 -> defer func -> return r
}
