package main

import (
	"fmt"
	"time"
)

func main() {
	//方法1
	timer1 := time.NewTimer(time.Second * 2) //创建一个两秒的定时器
	fmt.Printf("time.Now:%v\n", time.Now())
	t := <-timer1.C //这个地方的C是一个通道，会阻塞等待相应的时间
	fmt.Printf("t:%v\n", t)
	fmt.Println("-----------------------------")
	//方法2
	timer2 := time.NewTimer(time.Second * 2) //创建一个两秒的定时器
	fmt.Printf("time.Now:%v\n", time.Now())
	<-timer2.C
	fmt.Printf("2s后 time.Now:%v\n", time.Now())
	fmt.Println("-----------------------------")
	//方法3
	fmt.Printf("time.Now:%v\n", time.Now())
	time.Sleep(time.Second * 2)
	fmt.Printf("2s后 time.Now:%v\n", time.Now())
	fmt.Println("-----------------------------")
	//方法4
	fmt.Printf("time.Now:%v\n", time.Now())
	<-time.After(time.Second * 2)
	fmt.Printf("2s后 time.Now:%v\n", time.Now())
	fmt.Println("-----------------------------")

	//停止计时器
	timer3 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer3.C //被停止后，后面的句子也不会执行了
		fmt.Println("Timer3 expired")
	}()
	stop := timer3.Stop()
	if stop {
		fmt.Println("Timer3 stop")
	}
	time.Sleep(time.Second * 5)
	fmt.Println("-----------------------------")

	//重新设置计时器的值
	timer4 := time.NewTimer(time.Second * 5)
	timer4.Reset(time.Second * 1)
	<-timer4.C
	fmt.Println("end-------------------------")
}
