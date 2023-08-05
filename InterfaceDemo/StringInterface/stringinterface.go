package main

import "fmt"

type Demo struct {
	str string
}

type Demo1 struct {
	str string
}

func (d Demo) String() (sout string) {
	sout = "我是方法String"
	return
}

type Element interface{}
type List []Element

func main() {
	x := Demo{"1242353"}
	x1 := Demo1{"1242353"}
	fmt.Println("aaa", x)
	fmt.Println("aaa", x1)

}
