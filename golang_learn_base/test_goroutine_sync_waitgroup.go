package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func showMsg(i int) {
	defer wg.Done() //相当于wg.Add(-1)
	fmt.Printf("i: %v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //每启动一个协程就加1
		go showMsg(i)
	}

	wg.Wait() //此时主协程会一直阻塞等待所有其他协程执行完毕，达到同步的效果
	//主协程
	fmt.Println("end...")
}
