package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sprint(s string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("s: %v\n", s)
		time.Sleep((time.Millisecond * 10))
	}
}

var values = make(chan int) //创建了一个无缓冲通道

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10) //生成一个介于0到9（包括0和9）之间的随机整数
	fmt.Printf("value: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value //将随机生成的值传入无缓冲通道
}

// 在任务函数前面添加一个关键字go就可以创建协程
// go提供了一种称为通道的机制，用于在goroutine（协程）之间共享数据，通道在声明时需要指明数据类型，
// 数据在通道上传递时一次只有一个协程可以访问数据项，因此不会发生数据竞争
// 有两种通道：缓冲通道和无缓冲通道，前者用于异步通信，后者用于同步通信（无缓冲通道保证在发送和接收瞬间执行两个goroutine之间的交换）
// 对同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的，发送和接收操作对于元素值的处理是原子的，发送操作在完全完成之前会被阻塞，接收操作也是如此
func main() {
	fmt.Println("正常执行，顺序执行")
	sprint("tom")
	sprint("jimmy")
	fmt.Println("---------------------")
	fmt.Println("使用协程，并发执行")
	go sprint("tom") //在这里会创建一个协程，并发执行。main函数中本来就有一个主协程用来顺序执行main函数中的语句
	sprint("jimmy")
	fmt.Println("---------------------")
	defer close(values) //通道在程序结束前需要关闭
	var v int
	go send()
	fmt.Println("waiting...")
	v = <-values //若没有数据会在这里阻塞
	fmt.Printf("v: %v\n", v)
}
