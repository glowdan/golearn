package main

import "fmt"
type composer struct {
	name string
	birth int
}
func main() {
	antonio := composer{"antonio", 99}
	agnes := new(composer)
	agnes.name, agnes.birth = "anges", 1876
	julia := &composer{}
	julia.name, julia.birth = "julia", 2000
	angusta := &composer{"angusta", 2014}
	fmt.Println(antonio)
	fmt.Println(agnes)
	fmt.Println(julia)
	fmt.Println(angusta)
	arrayTmp()
	sliceTmp()
	mapTmp()
}

func arrayTmp() {
	var buffer [20]byte
	var grid1 [3][3]int
	grid1[1][0], grid1[1][1], grid1[1][2] = 2,4,8
	grid2 := [3][3]int{{4,3}, {8,3,3}}
	cities := [...]string{"a", "b", "c"}
	cities[len(cities)-1] = "d"
	fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
	fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
	fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)
	fmt.Printf("%-8T %2d %v\n", cities, len(cities), cities)
}

func sliceTmp() {
	s := []string{"a", "b", "c", "d", "e", "f", "g"}
	t := s[:5]
	u := s[3:len(s)-1]
	fmt.Println(s, t, u)
	u[1] = "X"
	fmt.Println(s, t, u)
	makeslice := make([]int, 20, 20)
	for i:= range makeslice{
		makeslice[i] = i
	}
	makeslice = append(makeslice, 3)
	makeslice = append(makeslice, 3)
	makeslice = append(makeslice, 3)
	makeslice = append(makeslice, 3)
	makeslice = append(makeslice, 3)
	fmt.Println(makeslice)
	fmt.Println(cap(makeslice))
}

func insertStringSliceCopy(slice, insertion[]string, index int) []string {
	result := make([]string, len(insertion)+len(slice))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func mapTmp() {
	mapT := map[string]int{}
	mapT["hell"] = 0
	fmt.Println(mapT)
	if value,found := mapT["hell"];found {
		fmt.Printf("found:%d", value)
	} else {
		fmt.Printf("not found%d", value)
	}
}
