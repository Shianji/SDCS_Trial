package main

import (
	"fmt"
	"os"
	"time"
)

// 创建文件
func createfile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func makedir() {
	err := os.Mkdir("testdir", os.ModePerm) //os.ModePerm代表权限"777"
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	err2 := os.MkdirAll("a/b/c", os.ModePerm) //创建层级目录
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 删除目录或者文件
func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.RemoveAll("a") //删除整个目录
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 获取和修改当前工作目录
func wd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("/home/adduser/workspace/homework")
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)
}

// 其他方法如：重命名文件os.Rename(原文件名，新文件名)、
// 读文件和写文件
func read_write() {
	_, err := os.Create("b.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	os.WriteFile("b.txt", []byte("hello"), os.ModePerm) //写入,go语言中的os.WriteFile和os.ReadFile方法在读写文件时无需显示地打开或关闭文件
	b, _ := os.ReadFile("b.txt")
	fmt.Printf("file content: %v\n", string(b))
	os.Remove("b.txt")
}

func main() {
	createfile()
	makedir()
	time.Sleep(time.Second * 3)
	remove()
	wd()
	fmt.Println("-----------------------------")
	read_write()
}
