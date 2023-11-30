package main

import "fmt"

// 结构体方法和属性是分开的
type person struct {
	name string
	age  int
}

// (per person)是接受者，说明eat函数是属于person结构体的，接受者还可以是结构体指针类型
// 接受者不一定非要是结构体，还可以是slice、map、channel、func等
func (per person) eat() {
	fmt.Printf(per.name + "is eating\n")
}

func (per person) sleep() {
	fmt.Printf(per.name + "is sleeping\n")
}

func main() {
	var tom = person{
		"tommy",
		28,
	}
	tom.sleep()
	tom.eat()
}
