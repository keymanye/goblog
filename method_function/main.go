package main

import "fmt"

type Ractangle struct {
	width, height float64
}

func area(r Ractangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Ractangle{10, 20}
	r2 := Ractangle{20, 30}
	fmt.Println("Area of r1 is ", area(r1))
	fmt.Println("Area of r2 is ", area(r2))
}
