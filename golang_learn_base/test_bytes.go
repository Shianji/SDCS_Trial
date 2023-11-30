package main

import (
	"bytes"
	"fmt"
	"strings"
)

func trans() {
	var a int = 100
	var b byte = 10
	a = int(b) //强制类型转换
	a = 100
	b = byte(a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	var s string = "hello world"
	var c = []byte{97, 98, 99}
	s = string(c)
	c = []byte(s)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("c: %v\n", c)
}

func contain() {
	s := "hello_world.com"
	b1 := []byte(s)
	b2 := []byte("hello_world")
	b3 := []byte("Hello_world")
	i := bytes.Contains(b1, b2) //该函数判断b1中是否包含b2
	fmt.Printf("i: %v\n", i)
	i = bytes.Contains(b1, b3)
	fmt.Printf("i: %v\n", i)

	i = strings.Contains("hello world", "hello") //字符串strings库中也有类似的方法
	fmt.Printf("i: %v\n", i)
}

func buf() {
	var b bytes.Buffer
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %v\n", b)
	fmt.Println("-----------------------------")
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1: %T\n", b1)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b1: %v\n", *b1)
	fmt.Println("-----------------------------")
	var b2 = bytes.NewBuffer([]byte("hello world"))
	fmt.Printf("b2: %T\n", b2)
	fmt.Printf("b2: %v\n", b2) //b2指向的地址刚好是b2结构体的第一项也就是字符串的首地址，所以输出了hello world
	fmt.Printf("b2: %v\n", *b2)
	fmt.Println("-----------------------------")
	n, _ := b.WriteString("hello_world!") //其他写入方式 WriteByte、WriteRune等与此类似
	fmt.Printf("n: %v\n", n)
	fmt.Printf("b: %s\n", string(b.Bytes())) //bytes.Buffer类型的内部实现结构体是不公开的，这意味着其内部成员对用户是不可见的（bytes.Buffer结构体成员buf和off的首字母是小写的，表示它们是不可导出的）,所以只能通过b.Bytes()访问
	fmt.Println("-----------------------------")
	bf := make([]byte, 2)
	n, _ = b1.Read(bf) //其他读取方式  ReadByte、ReadRune等与此类似
	fmt.Printf("n: %v\n", n)
	fmt.Printf("bf: %v\n", string(bf))
}

func main() {
	trans()
	fmt.Println("-----------------------------")
	contain()
	fmt.Println("-----------------------------")
	//其他一些函数如Count计算子串出现次数，Repeat重复输出，Replace子串替换，Runes可将字符转换为unicode切片(一个汉字占3个字符)，Join用来连接切片等等
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune

	fmt.Printf("ba%s\n", bytes.Repeat([]byte("na"), 3))

	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1)) //-1表示有几个替换几个

	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Printf("转换前字符串长度len(s): %v\n", len(s))
	fmt.Printf("转换后字符串长度len(r): %v\n", len(r))

	s1 := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s1, []byte(", ")))
	fmt.Println("-----------------------------")
	buf()
}
