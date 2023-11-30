package main

import "fmt"

func main() {
	s := []int{1, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)
	s1 := []int{4, 5, 6}
	i1 := append(s, s1...) //后面的...表示解包
	fmt.Printf("i1: %v\n", i1)

	fmt.Println("--------------------------------")
	s2 := "hello world"
	fmt.Printf("len(s2): %v\n", len(s2)) //len可以用来获取切片、通道、字符串等等的长度

	fmt.Println("--------------------------------")
	name := "tom"
	age := 26
	print(name, " ", age, "\n")
	println(name, " ", age)

	// fmt.Println("--------------------------------")
	// defer fmt.Println("panic后我还是会执行")
	// panic("完蛋 发生异常了") //抛出异常，后面的代码部分不会再执行，前面的defer语句还是会执行
	// fmt.Println("end...............")

	//make只能初始化类型为slice、map、chan的数据；而new可以分配任意类型的数据
	//make返回的是引用本身，即T；new返回的是指针，即*T
	//make分配后，会被初始化；new分配的空间会被清零
	b := new(bool)
	fmt.Printf("*b: %v\n", *b)
	fmt.Printf("b: %T\n", b)
	j := new(int) //Go的垃圾收集器负责自动管理内存，因此在使用new分配内存时，不需要手动释放它
	fmt.Printf("*j: %v\n", *j)
	fmt.Printf("j: %T\n", j)
	ss := new(string)
	fmt.Printf("*ss: %v\n", *ss)
	fmt.Printf("ss: %T\n", ss)
	fmt.Println("-------------------------")
	var p *[]int = new([]int)
	var v []int = make([]int, 10)
	fmt.Printf("p: %T\n", p)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("v: %T\n", v)
	fmt.Printf("v: %v\n", v)
}
