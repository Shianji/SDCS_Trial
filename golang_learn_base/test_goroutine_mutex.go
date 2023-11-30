package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wg sync.WaitGroup //sync.WaitGroup是一个结构体
var lock sync.Mutex   //互斥锁

func add() {
	time.Sleep(time.Millisecond * 2)
	lock.Lock() // 不加锁结果不可预测，有时候会出错时
	i++         //不是原子操作
	fmt.Printf("i++, i=%v\n", i)
	lock.Unlock()
	wg.Done()
}

func sub() {
	time.Sleep(time.Millisecond * 2)
	lock.Lock()
	i-- //不是原子操作
	fmt.Printf("i--, i=%v\n", i)
	lock.Unlock()
	wg.Done()
}

func main() {
	// for i := 0; i < 50; i++ { //都是在主协程中运行，不存在同步的概念，i最后肯定是初值100
	// 	add()
	// 	sub()
	// }
	// fmt.Println("--------------------------------")
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Printf("end i: %v------------------------\n", i)
}
