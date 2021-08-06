package main

import (
	"fmt"
)

type Employee struct {
	name   string
	age    int
	salary float32
}

func (e Employee) calculateSalaryPercentage() float32 {
	return e.salary * 0.1
}

func (e Employee) getName() string {
	return e.name
}

//asi no funciona, debo colocar un apuntador
func (e Employee) setName(name string) {
	e.name = name
}

func (e *Employee) setNamePointer(name string) {
	e.name = name
}

func main() {

	diego := Employee{"diego", 31, 8000000}
	fmt.Println(diego)
	diego.setName("Armando")
	fmt.Println(diego)
	diego.setNamePointer("Armando")
	fmt.Println(diego)
}
