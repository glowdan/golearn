package main

import (
	"fmt"
	"runtime"
	"time"
)

var fun = func(x, y int)int{
	return x+y
}

func addr() func(int) int {
	sum := 0
	y := 0
	return func(x int) int{
		//fmt.Println("x", y)
		sum = x + y
		y = sum
		//fmt.Println("y", y)
		fmt.Println(sum)
		return sum
	}
}


func main() {
	function := fun
	fmt.Println(function(1, 2))
	sarha, darlin := addr(), addr()
	sarha (2)
	sarha (2)
	for i:=1;i<10;i++ {
		darlin(i)
	}
	switch os := runtime.GOOS;os{
	case "darwin":
		fmt.Println("osx")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("unknown")
	}
	today := time.Now().Weekday();
	fmt.Println(today)
	fmt.Println(time.Saturday)
	switch time.Saturday{
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 6:
		fmt.Println("Yesterday")
	default:
		fmt.Println("Too far away.")
	}
	switch {
	case time.Now().Hour()<12:
		fmt.Println("am")
	case time.Now().Hour()>11:
		fmt.Println("pm")
	}
}
