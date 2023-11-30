package main

import (
	"fmt"
	"log"
	"os"
)

func logp() {
	log.Print("my log")
	log.Printf("my log %d", 100) //即格式化输出

	name := "tom"
	age := 26
	log.Println(name, " ", age)
}

func logpanic() {
	fmt.Println("--------------------------")
	defer fmt.Println("log.Panic 结束后再执行......")
	log.Panic("my panic") //log.Panic会抛出异常，但是在log.Panic前面的defer语句仍然会执行
	fmt.Println("END.........")
}

func logfatal() {
	fmt.Println("--------------------------")
	defer fmt.Println("log.Fatal 结束后再执行......")
	log.Fatal("my fatal") //log.Fatal不会抛出异常但是前面的defer语句不会再执行，它的实际实现调用了os.Exit(1)
	fmt.Println("END.........")
}

func logtofile() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //这个函数用来设置日志的输出格式，如log.Lshortfile表示以短文件形式即相对路径输出
	log.SetPrefix("Mylog:")                              //设置输出的日志前缀
	f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0664)
	defer f.Close()
	if err != nil {
		log.Fatal("日志文件错误")
	}
	log.SetOutput(f) //设置将日志输出到文件a.txt中
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.Print("my log")
}

var logger *log.Logger

func main() {
	// logp()
	// logpanic()
	// logfatal()
	// logtofile()
	f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0664)
	defer f.Close()
	if err != nil {
		log.Fatal("日志文件错误")
	}
	logger = log.New(f, "Mylog:", log.Ldate|log.Ltime|log.Lshortfile) //也可以直接通过这个函数设置日志属性,相当于将前面的设置全部写到了一起
	logger.Print("my log")
}
