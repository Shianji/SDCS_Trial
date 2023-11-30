package main

import "fmt"

func main() {
	//使用append方法可以实现切片的添加
	var s1 []int
	fmt.Printf("s1: %v\n", s1)
	s1 = append(s1, 100)
	fmt.Printf("s1: %v\n", s1)
	s1 = append(s1, 200)
	fmt.Printf("s1: %v\n", s1)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
	s1 = append(s1, 400)
	fmt.Printf("s1: %v\n", s1)
	//使用append方法可以实现切片的删除,其中...是展开符
	s1 = append(s1[:1], s1[2:]...)
	fmt.Printf("s1: %v\n", s1)
	//可以直接使用索引修改切片的元素
	s1[1] = 249
	fmt.Printf("s1: %v\n", s1)
	//切片的复制，若是直接赋值则两个切片会指向同一片内存地址，修改其中一个另外一个也会随之改变
	s2 := []int{1, 2, 3, 4, 5, 6}
	s3 := s2
	//如果使用copy(dst,src)函数，就会指向不同的内存，修改其中一个的值另外一个不受影响
	var s4 = make([]int, 6)
	copy(s4, s2)
	s2[2] = 50
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)
}
