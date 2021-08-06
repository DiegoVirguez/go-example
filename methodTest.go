package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//si lo coloco en los argumentos se genera un error
//type Vertex has no field or method Abs
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
