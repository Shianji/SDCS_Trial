package main

import (
	"fmt"
	"runtime"
	"time"
)

func showString(s string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("s: %v\n", s)
	}
}

func showInt(j int) {
	for i := 0; i < 8; i++ {
		fmt.Printf("j: %v\n", j)
		if i == 4 {
			runtime.Goexit() //用于退出当前协程
		}
	}
}

func show1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("这是函数show1,第%v次输出\n", i)
		// time.Sleep(time.Millisecond * 100)
	}
}

func show2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("这是函数show2,第%v次输出\n", i)
		// time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showString("java")
	for j := 0; j < 3; j++ {
		runtime.Gosched() //这句代码的作用是主协程将cpu让出来给其他协程执行，相当于调度。如果不加，可能其他协程还没执行，主协程就执行完然后程序退出了
		fmt.Println("python")
	}
	fmt.Println("--------------------")
	go showInt(40)
	time.Sleep(time.Second)
	fmt.Println("--------------------")
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) //查看当前CPU核心数
	//为什么不修改可用核心数，还是顺序输出，不应该交替输出吗（go不是默认所有可用核心都能使用吗）(有可能是运行太快了)
	go show1()
	go show2()
	time.Sleep(time.Second)
	fmt.Println("--------------------")
	runtime.GOMAXPROCS(1) //设置使用的核心数
	go show1()
	go show2()
	fmt.Println("end..........................")
	time.Sleep(time.Second)
}
