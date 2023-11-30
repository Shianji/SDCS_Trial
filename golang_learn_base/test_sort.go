package main

import (
	"fmt"
	"sort"
)

// sort库中的排序主要是通过以下接口来实现的，我们可以自定义数据类型，只要实现了下面接口中的方法，就可以使用该库中的函数func sort.Sort(data Interface)来排序
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }
//sort库中的那些排序函数也是通过上面的方式来实现的，库中实现了int、float64和string类型的

// 自定义一个数据类型如下：
type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	s := []int{2, 1, 4, 6, 3}

	fmt.Printf("排序前s: %v\n", s)
	sort.Ints(s)
	fmt.Printf("排序后s: %v\n", s)

	fmt.Println("下面是自定义的数据类型NewInts:")
	u := NewInts{1, 8, 6, 3, 7, 4}
	fmt.Printf("u: %v\n", u)
	sort.Sort(u)
	fmt.Printf("u: %v\n", u)
}
