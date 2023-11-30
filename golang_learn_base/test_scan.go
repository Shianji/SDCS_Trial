package main

import "fmt"

func main() {
	var name string
	var age int
	var email string

	fmt.Println("请输入姓名，年龄，邮箱，用空格分隔")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
}
