package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	rand.Seed(2)
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Nextafter(2,3))
	fmt.Println(math.Pi)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(3*3 + 4*4))
	var z int = int(f)
	fmt.Println(x, y, z)
}
