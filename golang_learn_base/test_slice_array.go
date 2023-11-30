package main

import "fmt"

func main() {
	//len可以输出数组、切片等的长度;cap用来输出切片的容量
	var num = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("length:%d, num: %v\n", len(num), num)

	//make可以用来创建切片
	var slice1 []int
	//make([]T, length, capacity),length 是切片的长度，表示切片中当前的元素个数;capacity 是切片的容量，
	//表示切片底层数组的大小，也就是切片可以扩展的最大长度。如果不提供容量参数，切片的容量将与长度相同
	var slice2 = make([]int, 2, 100)
	//切片可以直接初始化
	slice3 := []int{1, 2, 3, 1, 5}

	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)
	fmt.Printf("len(slice1): %v\n", len(slice1))
	fmt.Printf("len(slice2): %v\n", len(slice2))
	fmt.Printf("len(slice3): %v\n", len(slice3))
	fmt.Printf("cap(slice1): %v\n", cap(slice1))
	fmt.Printf("cap(slice2): %v\n", cap(slice2))
	fmt.Printf("cap(slice3): %v\n", cap(slice3))
	fmt.Printf("slice3[3]: %v\n", slice3[3])

	//切片还可以通过其他切片或者数组来初始化，[a:b]取出的切片包含左侧的元素a而不包含右侧的元素b
	// s := []int{1, 5, 6, 7, 1, 2, 3, 6}
	s := [...]int{1, 5, 6, 7, 1, 2, 3, 6}
	fmt.Printf("s: %T\n", s) //s的类型为[8]int是数组
	s1 := s[:]
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1: %T\n", s1) //s1的类型为[]int是切片
	s2 := s[2:]
	fmt.Printf("s2: %v\n", s2)
	s3 := s[:5]
	fmt.Printf("s3: %v\n", s3)
	s4 := s[4:7]
	fmt.Printf("s4: %v\n", s4)

}
