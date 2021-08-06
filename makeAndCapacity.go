package main

import "fmt"

func main() {
	slice := make([]int, 3)
	slice = append(slice, 2)
	fmt.Println(slice)
}
