package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//常见格式化输出占位符%v表示变量  %T表示输出变量的类型  %p表示指针  %s表示字符串 其他还有%d、%f等等
	//常见转义字符：\n换行 \t制表符 \'单引号 \"双引号 \\反斜杠
	var s1 string = "hello world!1"
	var s2 = "hello world!2"
	s3 := "hello world!3"
	//反引号可以用来定义多行字符串
	s4 := `
	line1
	line2
	line3
	`
	//字符串连接可以直接使用+
	s5 := s1 + s2
	//格式化打印字符串
	s6 := fmt.Sprintf("S1=%v,s2=%s", s1, s2)
	//使用Join方法进行连接
	s7 := strings.Join([]string{s2, s3}, ",")
	//通过写入buff连接字符串
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(",")
	buffer.WriteString("shianji")

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)
	fmt.Printf("s5: %v\n", s5)
	fmt.Printf("s6: %v\n", s6)
	fmt.Printf("s7: %v\n", s7)
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	//字符串切片
	s := "hello world"
	a := 2
	b := 5
	fmt.Printf("s[a]:%c\n", s[a])
	fmt.Printf("s[a,]:%v\n", s[a:])
	fmt.Printf("s[,b]:%v\n", s[:b])
	fmt.Printf("s[,]:%v\n", s[:])

	fmt.Printf("shianji %s\n", "like learning")
}
