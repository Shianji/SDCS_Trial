package main

import "fmt"

func f() {

}

func main() {
	name := "shianji"
	age := 26
	//不能用0或非0来表示true或者false，否则会报错
	m := true
	a := 100
	p := &a
	//数组,...表示省略数组长度
	num := [...]int{1, 2, 3}
	var balance = [2]float32{3.14, 4.13}
	//切片类型,可看成是动态数组，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
	split := []int{5, 6}
	//%T格式化输出数据类型
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", num)
	fmt.Printf("num: %v\n", num)
	fmt.Printf("balance: %v\n", balance)
	fmt.Printf("%T\n", balance)
	fmt.Printf("%T\n", split)
	fmt.Printf("%T\n", f)
}
