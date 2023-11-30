package main

import (
	"fmt"
	"io"
	"os"
)

func openfile() {
	f, err := os.Open("a.txt") //若不存在会返回错误，这个方法打开的文件是只读的
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}

	f1, err1 := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, 0755) //若不存在会创建一个新的文件
	if err1 != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f1.Name(): %v\n", f1.Name())
		f.Close()
	}
}

func write() {
	f, _ := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, 0755)
	f.Write([]byte("hello world helloooo wwwworld"))
	f.WriteString("WRITE STRING") //直接写入字符串
	f.Close()
}

func read() {
	f, _ := os.Open("b.txt")
	for {
		buf := make([]byte, 10)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("file c.txt: %v\n", string(buf))
	}
	f.Close()
}

// 其他方法如f.ReadAt(缓冲区buf，偏移量lseek)可以设置从某个偏移量开始读取，还有f.WriteAt可设置写入位置；os.ReadDir方法可以用来遍历目录；f.Seek(偏移量，何处whence)方法可以设置文件偏移量
func main() {
	openfile()
	write()
	read()
}
