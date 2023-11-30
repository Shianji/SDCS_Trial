package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanString = make(chan string)

func main() {
	go func() {
		chanInt <- 100 //这个协程执行到这里会阻塞，它一直等待知道chanInt通道中的数据被读走之后才会继续向下执行
		chanString <- "hello"
		close(chanInt) //如果不关闭则下面的select会一直进入default分支
		close(chanString)
		fmt.Println("通道函数退出")
	}()

	for {
		select { //select语句用于处理异步IO操作，select会监听case语句中的channel的读写操作（select语句中的case语句必须是一个channel操作）当这些channel可以读写时，将会出发相应的操作。
		//select语句中的default语句总是可以运行的，如果有多个case语句满足条件可执行，则select会随机选择其中一个执行，若没有case语句满足条件则执行default语句
		case i := <-chanInt:
			fmt.Printf("i: %v\n", i)
		case s := <-chanString:
			fmt.Printf("s: %v\n", s)
		default:
			fmt.Println("default..........")
		}
		time.Sleep(time.Second)
	}
}
