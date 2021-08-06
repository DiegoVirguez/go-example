package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

//una estructura es una coleccion de fields
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
