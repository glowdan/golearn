package main

import (
	"fmt"
	"math"
)
type Vegas struct {
	va int
	vb float64
}

func main() {
	sum := 0
	for i:=0;i<10;i++{
		sum += i
		fmt.Println(sum)
	}
	sum2 := 0
	for ;sum2<1000;sum2 += 100 {
		fmt.Println(sum2)
	}
	fmt.Println(sum2)
	sum3 := 0
	for sum3<1000 {
		sum3 += 500
	}
	fmt.Println("sss", sum3)
	if sum>8 {
		fmt.Println(22)
	}
	a := func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
		return v
		}
		return lim
	}
	fmt.Println(a (1, 3,5))
	sq := 1.0
	rt := 2.0
	for i:=0;i<10;i++{
		sq = sq - (sq*sq-rt)/(2*sq)
		fmt.Println(sq);
	}
	fmt.Println(Vegas{1, 1.1})
	vv := Vegas{2, 2.2}
	vs := &vv
	vn := Vegas{vb:9.9}
	vs.va=3
	fmt.Println(vs)
	fmt.Println(vv)
	fmt.Println(vn)
	vm := new(Vegas)
	vm.va = 8
	vm.vb = 9
	fmt.Println(vm)
}

