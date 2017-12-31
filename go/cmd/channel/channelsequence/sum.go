package main

import (
	"fmt"
)

func sum(a []int, c chan<- int) {
	var result int
	for _, v := range a {
		result += v
	}
	c <- result
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	fmt.Println("sum of a is:", <-c+<-c)
}
