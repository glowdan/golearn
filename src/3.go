package main

import "fmt"
type vex struct {
	a int
	b int
}
func main() {
	var arr [2] string
	arr[0] = "9"
	arr[1] = "2"
	fmt.Println(arr)
	slice := []int{0, 2, 4, 8}

	fmt.Println(slice[:])
	slice2 := make([]int, 3, 5)
	fmt.Printf("%d, %d,  %v", len(slice2), cap(slice2), slice2)
	var sliceEmpty  []int
	fmt.Printf("%d, %d,  %v", len(sliceEmpty), cap(sliceEmpty), sliceEmpty)
	if nil==sliceEmpty {
		fmt.Println(999)
	}
	for _,ou := range slice {
		fmt.Println(ou)
	}
	//-----------------------------
	m := make(map [string]int)
	m["s"] = 1
	fmt.Println(m)
	fmt.Println(len(m))
//	var mm = map [string]vex{
//		"s": vex{1,1},
//	}
	mmm := make(map [string] vex )
	mmm ["ss"] = vex{2,2}
	fmt.Println(mmm)
}
