package main

import "fmt"

// go中的结构体也可以嵌套使用，向C语言一样
// 相当于定义了一个新类型，结构体定义时成员后面是不加逗号的
type custmomer struct {
	name  string
	age   int
	email string
}

type Person struct {
	no   int
	name string
}

func showPerson(person Person) {
	person.name = "lisi"
	person.no = 2
	fmt.Printf("person: %v\n", person)
}

func main() {
	var c1 custmomer
	//未初始化时结构体成员相应地被初始化为0或者nil
	fmt.Printf("c1: %v\n", c1)
	c1.age = 26
	c1.name = "shianji"
	c1.email = "shianji@qq.com"
	fmt.Printf("c1: %v\n", c1)

	c2 := custmomer{}
	fmt.Printf("c2: %v\n", c2)
	fmt.Println("--------------------")
	//按找定义顺序进行结构体初始化,注意后面的逗号(不能和键值对初始化混合使用)
	var c3 custmomer
	c3 = custmomer{
		"zhangsan",
		18,
		"zhangsan@qq.com",
	}
	fmt.Printf("c3: %v\n", c3)
	fmt.Println("--------------------")
	//按照键值对部分初始化
	c4 := custmomer{
		name: "lisi",
		age:  16,
	}
	fmt.Printf("c4: %v\n", c4)
	fmt.Println("--------------------")
	//结构体指针
	var pc *custmomer
	pc = &c3
	fmt.Printf("pc: %p\n", pc)
	fmt.Printf("*pc: %v\n", *pc)
	fmt.Printf("pc.name: %v\n", pc.name) //注意通过指针来访问结构体成员的方式，直接通过.就可以访问
	fmt.Printf("pc.age: %v\n", (*pc).age)
	fmt.Println("--------------------")
	//还可以通过new来创建指针，它用于分配内存，并返回一个指向该内存的指针
	//在go中，使用new函数创建的内存不需要手动回收，因为Go具有自动内存管理机制，也就是垃圾回收
	var pc2 = new(custmomer)
	fmt.Printf("pc2: %v\n", pc2)
	fmt.Printf("*pc2: %v\n", *pc2)
	pc2.age = 11
	pc2.email = "wangwu@qq.com"
	pc2.name = "wangwu"
	fmt.Printf("pc2: %v\n", pc2)
	fmt.Printf("*pc2: %v\n", *pc2)
	fmt.Println("--------------------")
	//定义一个匿名结构体
	var tom struct {
		name        string
		age, weight int
	}
	tom.age = 28
	tom.name = "tommy"
	tom.weight = 65
	fmt.Printf("tom.age: %v\n", tom.age)
	fmt.Printf("tom.name: %v\n", tom.name)
	fmt.Printf("tom.weight: %vkg\n", tom.weight)
	fmt.Println("--------------------")
	//结构体也可以作为函数参数传递，但是仍然是值传递，对形参的修改不改变实参的内容，但如果传递结构体指针就会有影响
	var person = Person{
		1,
		"zhangsan",
	}
	fmt.Printf("before call,person: %v\n", person)
	showPerson(person)
	fmt.Printf("after call,person: %v\n", person)
}
