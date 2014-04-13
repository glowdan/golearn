package main

import (
	"math"
	"fmt"
	"os"
)

type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
type myfloat float64
type abser interface {
	Abs() float64
}
func (mf myfloat) abb () float64{
	return float64(mf+1.0)
}
func (mf myfloat) Abs () float64{
	return float64(mf+1.0)
}
func (v *Vertex) scale(z float64) {
	v.X *= z
	v.Y *= z
}


type Reader interface {
	Read(b []byte) (n int, err error)
}
type Writer interface {
	Write(b []byte) (n int, err error)
}
type ReaderWriter interface {
	Reader
	Writer
}
func main() {
	v := Vertex{3, 4}
	v.scale(5)
	fmt.Println(v.Abs())
	my := myfloat(3.0)
	fmt.Println(my.abb())
	var a abser
	f := myfloat(6.4)
	vv := Vertex{3, 4}
	a = f
	a = &vv
	fmt.Println(a.Abs())
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, world\n")
}
