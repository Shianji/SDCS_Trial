package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// //通过加锁实现同步
// var i = 100
// var lock sync.Mutex

// func add() {
// 	lock.Lock()
// 	i++
// 	lock.Unlock()
// }

// func sub() {
// 	lock.Lock()
// 	i--
// 	lock.Unlock()
// }

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1) //使用atomic中的原子变量的原子操作（cas:compare and swap）
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func read_write() {
	fmt.Println("---------------------")
	var j int32 = 200
	atomic.LoadInt32(&j) //read
	fmt.Printf("j: %v\n", j)
	atomic.StoreInt32(&j, 150) //write
	fmt.Printf("j: %v\n", j)
}

func cas() {
	fmt.Println("---------------------")
	var k int32 = 100
	r := atomic.CompareAndSwapInt32(&k, 100, 200) //比较后再修改，第二个参数是之前的值(先与之比较，若不相等说明值被其他协程修改了，返回False)，第三个参数是要修改的目标值
	fmt.Printf("r: %v\n", r)
	fmt.Printf("k: %v\n", k)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 2)
	fmt.Printf("i=%v\n", i)

	read_write()
	cas()
}
