package main

import (
	"fmt"
	"math"
)
type BitFlag int
func main() {
	count, _ := fmt.Printf("Hello world!")// 字符_ 表示占位符，此值不取
	fmt.Println(count);
	const ( //常量声明使用括号可以多个声明
		apple = iota//使用iota可以递增取值从0开始
		orange
		pear
	)
	const (
		banana = iota //重新从0开始
	)
	fmt.Println(banana)

	const (
		Active BitFlag = 1 << iota //1<<0
		Send                       //1<<1
 		Receive                    //1<<2
	)
	flag := Active | Send | Receive
	fmt.Println(flag)
	fmt.Println(math.E)
	fmt.Println(math.Exp(math.Pi))
	//复数运算
	f := 3.2e5
	x := -7.3 - 8.9i
	y := complex64(-18.3+8.91i)
	z := complex(f, 13.2)
	fmt.Println(x, real(y), imag(z))
}

func (flag BitFlag) String() string {//String() 这种类型的直接按照此方法进行输出
	return fmt.Sprintf("%b", flag)
}
