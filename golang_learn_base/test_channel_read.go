package main

import "fmt"

var channel = make(chan int)

func main() {
	go func() { //用匿名函数创建一个协程
		for i := 0; i < 7; i++ {
			channel <- (i + 2)
		}
		close(channel) //如果通道不再写入数据，即写端关闭，此时若读完通道中的数据后再读，会报死锁deadlock错误，如果关闭后再读则会读到0(int)或nil(string)
	}()

	var j int
	for k := 0; k < 4; k++ {
		j = <-channel
		fmt.Printf("j: %v\n", j)
	}

	for v := range channel { //也可以通过range来查看通道中剩余的内容，在本程序中协程每向无缓冲通道中写入一个数据这个for循环就输出一个，在这种遍历中，通道在后续没有数据写入时也要close，否则也会造成死锁
		fmt.Printf("v: %v\n", v)
	}

	//另外一种遍历通道的方式
	// 	for{
	// 		v,ok:=<-channel
	// 		if ok{
	// 			fmt.Printf("v: %v\n", v)
	// 		}else{
	// 			break
	// 		}
	// 	}
}
