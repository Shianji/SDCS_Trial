package main

import "fmt"

// 有一个返回值
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

// 有一个返回值，但不定义返回值
func sum1(a int, b int) int {
	return a + b
}

// 有多个返回值
func f1() (name string, age int) {
	name = "Tom"
	age = 20
	//return name, age
	return //不使用返回值名称效果是一样的
}

// 有多个返回值,return覆盖了返回值的名称，返回值名称没有被使用
func f2() (name string, age int) {
	a := "Tom"
	b := 20
	return a, b
}

func f() {
	fmt.Println("这个函数没有参数也没有返回值")
}

func f3(a []int) {
	a[0] = 100
}

// 还可以传递可变参数
func f4(args ...int) {
	for a, b := range args {
		fmt.Printf("a: %v b: %v\n", a, b)
	}
}

func f5(name string, age int, email string, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
	for a, b := range args {
		fmt.Printf("a: %v b: %v\n", a, b)
	}
}

// go的函数传参是值传递，和C语言一样(不一样的是传递数组名原数组内容不受影响)，但是如果参数是map、slice(切片)、interface、channel等类型，他们本身就是指针，所以修改值会有影响
func main() {
	r := sum(12, 20)
	fmt.Printf("r: %v\n", r)

	f()

	name, age := f1()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	a := []int{1, 2, 3}
	fmt.Printf("a: %v\n", a)
	fmt.Println("---------------")
	f3(a) //a的值被修改
	fmt.Printf("a: %v\n", a)
	fmt.Println("---------------")
	f4(1, 9, 6, 4)
	fmt.Println("---------------")
	f5("shianji", 26, "shianji@qq.com", 8, 6, 4, 9, 3)
}
