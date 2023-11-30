package main

import "fmt"

func main() {
	//类型声明
	var map_var map[string]string
	//使用make初始化
	map_var = make(map[string]string)
	//赋值
	map_var["name"] = "shianji"
	map_var["age"] = "26"
	map_var["email"] = "1662518187@qq.com"
	fmt.Printf("map_var: %v\n", map_var)

	//也可以在定义的时候初始化
	mapv := map[string]string{"name": "zhangsan",
		"age": "26", "email": "zhangsan@qq.com"}
	//map类型的输出是无序的
	fmt.Printf("mapv: %v\n", mapv)

	m2 := make(map[string]string)
	m2["name"] = "lisi"
	m2["age"] = "17"
	m2["email"] = "lisi@qq.com"
	fmt.Printf("m2: %v\n", m2)
	//判断某个键值对是否在map中,若存在返回相应的值和True，不存在则返回nil和False
	v, k := m2["name"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("k: %v\n", k)
	fmt.Println("---------------")
	v, k = m2["age1"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("k: %v\n", k)
	fmt.Println("---------------")

	//map的遍历,如果只给一个参数则只能遍历key，给两个参数则遍历key和value
	for k := range m2 {
		fmt.Printf("k: %v\n", k)
	}
	for k, v := range m2 {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

}
