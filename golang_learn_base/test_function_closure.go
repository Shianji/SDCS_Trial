package main

import "fmt"

// 函数闭包，
func add() func(a int) int {
	var x int
	return func(a int) int {
		x += a
		return x
	}
}

func main() {
	f := add()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))
	fmt.Println("----------------")
	f1 := add()
	fmt.Println(f1(1))
	fmt.Println(f1(2))
	fmt.Println(f1(3))

	fmt.Println(f(10))

}
