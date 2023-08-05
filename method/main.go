package main

import "fmt"

type Circle struct {
	radius float64
}

type Ractangle struct {
	width, height float64
}

//使用函数实现

func RactangleArea(r Ractangle) float64 {
	return r.width * r.height
}

func CircleArea(c Circle) float64 {
	return c.radius * 31415
}

// 使用方法实现
func (c Circle) area() float64 {
	return c.radius * 3.14
}

func (r Ractangle) area() float64 {
	return r.width * r.height
}

func main() {
	//函数调用
	r := Ractangle{10, 20}
	c := Circle{20}
	fmt.Println("Area of r1 is ", RactangleArea(r))
	fmt.Println("Area of r2 is ", CircleArea(c))
	r1 := Ractangle{width: 10, height: 20}
	r2 := Ractangle{width: 20, height: 30}

	c1 := Circle{radius: 20}
	c2 := Circle{radius: 30}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
