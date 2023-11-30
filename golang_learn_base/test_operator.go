package main

import "fmt"

func main() {
	a := 100
	b := 9

	fmt.Printf("a+b=%v\n", a+b)
	fmt.Printf("a-b=%v\n", a-b)
	fmt.Printf("a*b=%v\n", a*b)
	fmt.Printf("a/b=%v\n", a/b)

	//在go里面，自增自减不能被当做表达式赋值，只能作为单独的语句，例如 c:=a++是会报错的
	a++
	b--
	fmt.Printf("a+b=%v\n", a+b)

	r := a == b
	fmt.Printf("r: %v\n", r)
	r = a > b
	fmt.Printf("r: %v\n", r)
	r = a < b
	fmt.Printf("r: %v\n", r)
	r = a >= b
	fmt.Printf("r: %v\n", r)
	r = a <= b
	fmt.Printf("r: %v\n", r)
	r = a != b
	fmt.Printf("r: %v\n", r)
	//逻辑运算有与&&或||非！都和C语言一样，还有位逻辑、移位运算也一样
}
