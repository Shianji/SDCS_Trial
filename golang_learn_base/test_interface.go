package main

import "fmt"

// 在Go语言中，接口是一种抽象类型，它定义了一组方法的签名，但没有具体的实现。接口提供了一种方式，允许不同的类型来实现这些方法，从而实现多态性和代码的灵活性。
// 多个类型可以实现同一个接口，同一个类型也可以实现多个接口
type USBer interface {
	read()
	write()
}

type computer struct {
	name string
}

type mobile struct {
	size string
}

func (c computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("reading...")
}

func (c computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("writing...")
}

func (m mobile) read() {
	fmt.Printf("m.size: %v\n", m.size)
	fmt.Println("reading...")
}

func (m mobile) write() {
	fmt.Printf("m.size: %v\n", m.size)
	fmt.Println("writing...")
}

// go语言接收器实现的OCP原则可以扩展不能修改，是向上类型转换
func Read(usb USBer) {
	usb.read()
	usb.write()
}

// 接口定义时可以嵌套，即接口里面可以包含接口
type Flyer interface {
	fly()
}

type Swimer interface {
	swim()
}

type Fisher interface {
	Flyer
	Swimer
}

type fish struct {
}

func (f fish) swim() {
	fmt.Println("swim...")
}

func (f fish) fly() {
	fmt.Println("fly...")
}

func main() {
	c := computer{
		"computer1",
	}

	m := mobile{
		"5g",
	}
	//可以将实现了接口中所有的方法的结构体赋值给接口，若这些方法没有全部实现则会报错
	var usb USBer = c
	usb.read()
	usb.write()

	usb = m
	usb.read()
	usb.write()
	fmt.Println("-------------")

	Read(c)
	Read(m)

	fmt.Println("-------------")

	var f = fish{}
	var fi Fisher = f
	fi.fly()
	fi.swim()
}
