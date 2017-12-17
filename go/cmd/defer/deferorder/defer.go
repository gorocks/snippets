package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // ① go 没有延时求值, 此处会先计算 calc("10", a, b)
	a = 0
	defer calc("2", a, calc("20", a, b)) // ② 同理
	b = 1
	// 执行顺序: calc("10", a, b) -> calc("20", a, b) -> ② -> ①
}
