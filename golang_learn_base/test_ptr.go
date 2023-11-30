package main

import "fmt"

// go语言的指针不能进行偏移和运算，例如不能p++，go中的值类型（int、float、bool、string、array、struct）都有对应的指针类型
// 一个指针指向了一个值的内存地址，用&来取地址，*表示指针
func main() {
	var p *int
	fmt.Printf("p: %v\n", p) //未初始化的指针指向nil
	fmt.Printf("p_format: %T\n", p)

	var i int = 100
	p = &i
	fmt.Printf("i: %v\n", i)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("*p: %v\n", *p)
	fmt.Println("-------------")

	//以下相当于C语言中的指针数组
	var num = [3]int{1, 2, 3}
	var ptonum [3]*int
	for i := 0; i < len(ptonum); i++ {
		fmt.Printf("ptonum[i]: %v\n", ptonum[i])
	}
	fmt.Println("-------------")
	for i := 0; i < len(ptonum); i++ {
		ptonum[i] = &num[i]
		fmt.Printf("ptonum[i]: %v\n", ptonum[i])
	}
	fmt.Println("-------------")
	for i := 0; i < len(ptonum); i++ {
		fmt.Printf("*ptonum[i]: %v\n", *ptonum[i])
	}

	var pt *int = &num[0]
	fmt.Printf("%v\n", *pt)
	//下面这种用法是错误的
	// var ptr *int = &num
	// fmt.Printf("%v\n", *ptr)
}
