package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("i'm eating...")
}

func (a Animal) sleep() {
	fmt.Println("i'm sleeping...")
}

type Dog struct {
	a     Animal //可以理解为继承
	color string
}

type Cat struct {
	Animal //可以只写一个Animal
	color  string
}

func main() {
	dog := Dog{
		a: Animal{
			name: "dog1",
			age:  1,
		},
		color: "black",
	}
	dog.a.eat()
	dog.a.sleep()
	fmt.Printf("dog.color: %v\n", dog.color)
	fmt.Println("---------------")
	cat := Cat{
		Animal{
			name: "cat1",
			age:  2,
		},
		"white",
	}
	cat.eat()
	cat.sleep()
	fmt.Printf("cat.color: %v\n", cat.color)
}
