package main

import "fmt"

// 匿名变量的使用
func getName() (string, int, bool) {
	return "wangwu", 25, false
}

func getVar() {
	//初始化赋值
	/* var name string = "shianji"
	var age int = 26
	var m bool = true */

	//类型推断
	/* var name = "zhuwenhuan"
	var age = 26
	var m = true */

	//批量初始化
	// var name, age, m = "zhangsan", 25, true

	//短变量声明，只能用在函数内部
	/* 	name := "lisi"
	   	age := 25
	   	m := false */

	// name, age, m := getName()
	//go语言定义的变量必须使用，否则会报错，可以使用下划线来表示未使用的匿名变量
	name, age, _ := getName()

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	// fmt.Printf("m: %v\n", m)
}

func getConst() {
	//常量初始化时必须赋值，且之后不能再修改
	/* const PI float64 = 3.14
	const PI2 = 3.1415926
	const a, b, c = 1, 2, 3
	const (
		width  = 8
		height = 9
	)
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("PI2: %v\n", PI2)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("width: %v\n", width)
	fmt.Printf("height: %v\n", height) */

	//特殊的常量iota,可以被认为是一个可被编译器修改的常量，它默认开始值是0，每调用一次加1，遇到const关键字时被重置为0
	const (
		a1 = iota //0
		a2 = iota //1
		a3 = iota //2
	)

	const (
		X = iota // X的值是0，iota被重置
		Y        // Y的值是1
		Z        // Z的值是2
	)

	//可以使用_跳过某些值
	const (
		A = iota // 0
		_        // 跳过1
		B        // B的值是2
	)

	//中间插队的情况
	const (
		C = iota // 0
		D = 100
		E
		F
		G = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

	fmt.Printf("X: %v\n", X)
	fmt.Printf("Y: %v\n", Y)
	fmt.Printf("Z: %v\n", Z)

	fmt.Printf("A: %v\n", A)
	fmt.Printf("B: %v\n", B)

	fmt.Printf("C: %v\n", C)
	fmt.Printf("D: %v\n", D)
	fmt.Printf("E: %v\n", E)
	fmt.Printf("F: %v\n", F)
	fmt.Printf("G: %v\n", G)
}

func main() {
	getVar()

	getConst()

	fmt.Println("Hello, World!")
}
