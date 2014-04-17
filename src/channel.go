package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func  main () {
	a := []int{ 1, 2, 4, 3, 5, 6}
	c1 := make(chan int)
	c2 := make(chan int)
	go sum(a[:len(a)/2], c1)
	go sum(a[len(a)/2:], c2)
	x, y := <-c1, <-c2
	fmt.Println(x, y, x+y)
}
