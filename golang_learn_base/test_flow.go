package main

import "fmt"

func main() {
	var b int = 100
	//条件分支,go语言的大括号必须存在即使只有一条语句，并且左括号必须和if或else在同一行
	//go语言中，在if之后，条件语句之前可以使用变量初始化语句，但其作用域只在if语句内
	if a := 100; a > 50 {
		fmt.Println("a>50")
	} else {
		fmt.Println("a<=50")
	}
	//switch里面可以用break也可以不用，switch可以多条件匹配，不同于C语言go里面switch的case后面还可以是表达式
	switch b {
	case 100, 500:
		fmt.Println("b==100")
	case 200:
		fmt.Println("b==200")
	default:
		fmt.Println("b!=100")
	}
	//使用fallthrough可以向下穿透一行,即相执行多个case的话可以使用fallthrough
	fmt.Println("使用fallthrough可以向下穿透一行")
	switch b {
	case 100, 500:
		fmt.Println("b==100")
		fallthrough
	case 200:
		fmt.Println("b==200")
	case 300:
		fmt.Println("b==300")
	default:
		fmt.Println("b!=100")
	}

	//go里面只有for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	//for range循环
	x := [...]int{1, 5, 6, 4, 2, 9, 7}
	for i, v := range x {
		fmt.Println(i, v)
	}

	m := make(map[string]string, 0)
	m["name"] = "zhangsan"
	m["age"] = "20"
	m["email"] = "zhangsan@qq.com"
	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}

	s := "my name is shianji"
	for k, v := range s {
		fmt.Printf("%d:%c	", k, v)
		fmt.Println()
	}
}
