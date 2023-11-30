package main

import "fmt"

func sum(a int, b int) (ret int) {
	return a + b
}

func cmp(a int, b int) (ret int) {
	if a > b {
		return a
	} else {
		return b
	}
}

func sayHello(name string) {
	fmt.Printf("hello %v!\n", name)
}

// 可以将函数当做参数传递
func passFunc(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

// 函数也可以被当做返回值使用
func cacul(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	//type可以用来构造新类型的定义，函数也可以被看做是变量来进行赋值
	type fun func(int, int) int

	var s fun = sum
	fmt.Printf("1+2=%v\n", s(1, 2))
	var t fun = cmp
	fmt.Printf("1和2之间较大的是: %v\n", t(1, 2))
	fmt.Println("-----------------")
	passFunc("shianji", sayHello)
	fmt.Println("-----------------")
	cal := cacul("+")
	r := cal(2, 1)
	fmt.Printf("r: %v\n", r)
	cal = cacul("-")
	r = cal(2, 1)
	fmt.Printf("r: %v\n", r)
	fmt.Println("-----------------")
	//匿名函数的使用，go语言函数不能嵌套，但是在函数内部可以定义匿名函数
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	fmt.Printf("1和2之间较大的是 %v\n", max(1, 2))
	//也可以直接使用，自己调用自己
	func(a int, b int) {
		if a > b {
			fmt.Printf("a和b之间较大的值是 a=%v\n", a)
		} else {
			fmt.Printf("a和b之间较大的值是 b=%v\n", b)
		}
	}(1, 2)
	fmt.Println("------------")

	//类型定义，相当于创建了一个新的类型，原类型的方法新类型不能使用
	type Myint int
	var ty Myint = 199
	fmt.Printf("ty: %v\n", ty)
	fmt.Printf("ty_type:%T\n", ty)

	fmt.Println("------------")
	//类型别名，只是取了个别名，类型还是原来的类型，原来类型的方法，别名都可以使用
	type Myint1 = int
	var ty1 Myint1 = 199
	fmt.Printf("ty1: %v\n", ty1)
	fmt.Printf("ty1_type:%T\n", ty1)
}
