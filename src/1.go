package main

import (
	"fmt"
	"time"
	"os"
	"net"
	"math/rand"
	"math"
	"math/cmplx"
)
const (
	ca = 5
	Big = 1<<100
	Small = Big>>99
)
var c, python, java = true, false, "no!"
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	zz complex128 = cmplx.Sqrt(ca + 2i)
)

func main() {
	fmt.Println("Hello, world");
	fmt.Println("now time", time.Now())
	value,_ := os.Open("./1.txt")
	fmt.Println(value)
	netval, _ := net.Dial("tcp", "baidu.com:80")
	fmt.Println(netval)
	rand.Seed(time.Now().Unix())
	fmt.Println("random int", rand.Intn(7))
	fmt.Println("random int", rand.Intn(7))
	fmt.Println("random int", rand.Intn(7))
	fmt.Println("random int", rand.Intn(7))
	fmt.Println("random int", rand.Intn(7))
	fmt.Println("random int", rand.Intn(7))
	fmt.Println(math.Nextafter(3, 3))
	fmt.Println("pi ", math.Pi)
	plus, mini := add(1,3)
	fmt.Println(plus)
	fmt.Println(mini)
	fmt.Printf("1+3=%d %d", plus, mini)
	fmt.Println("2*6", cheng(2, 6))

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, zz, zz)
	fmt.Println(uint32(999))
	fmt.Println(float64(Big))
	fmt.Println(Small)
}

func add(x, y int) (int, int) {
	plus := x+y
	mini := y-x
	return plus, mini
}

func cheng(x, y int) (a int) {
	//z := x*y

	var z int
	z = 0
	a = y/x+z
	fmt.Println(c, python, java)
	return
}
