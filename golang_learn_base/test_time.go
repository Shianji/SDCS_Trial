package main

import (
	"fmt"
	"time"
)

func test1() {
	t := time.Now()
	fmt.Printf("t: %T\n", t)
	fmt.Printf("t: %v\n", t)
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T%T%T%T%T%T\n", year, month, day, hour, minute, second)
}

func test2() {
	//go语言里面的时间戳也是unix时间戳，是1970年1月1日至当前时间的总秒数，使用方法如下：
	now := time.Now()
	fmt.Printf("TimeStamp type: %T TimeStamp : %v\n", now.Unix(), now.Unix())
	//将时间戳转化为当前时间格式
	timestamp := time.Now().Unix()
	timeobj := time.Unix(timestamp, 0) //第二个参数0表示纳秒的偏移量。Unix时间戳通常以秒为单位，因此这里使用0表示不添加额外的纳秒偏移
	fmt.Printf("timeobj: %v\n", timeobj)
}

// 其他一些时间的操作如time.Add、time.Sub等见go官方文档
func main() {
	// test1()
	// test2()
	timenow := time.Now()
	fmt.Printf("timenow: %v\n", timenow.Format("2006/01/02 15:04")) //2006/01/02 15:04这个时间是go语言的诞生时间，修改时间输出格式时需要以此为模板
}
