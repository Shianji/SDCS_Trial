// Timer只执行一次，而Ticker可以周期的执行
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) //括号里面设置ticker的间隔时间

	// counter := 1
	// for _ = range ticker.C {
	// 	fmt.Println("ticker........")
	// 	if counter == 5 {
	// 		ticker.Stop()
	// 		break
	// 	}
	// 	counter++
	// }
	// fmt.Println("---------------------------")

	chanInt := make(chan int) //无缓冲通道
	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Printf("get num : %v\n", v)
		sum += v
		if sum >= 20 {
			fmt.Printf("sum is %v\n", sum)
			break
		}
	}
}
