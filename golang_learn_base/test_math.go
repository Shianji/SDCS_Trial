package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//一些数学常量的打印
	fmt.Printf("math.Pi: %v\n", math.Pi)
	fmt.Printf("math.E: %v\n", math.E)
	fmt.Printf("math.MaxInt: %v\n", math.MaxInt)
	fmt.Println("---------------------")
	//向上取整
	fmt.Printf("math.Ceil(math.Pi): %v\n", math.Ceil(math.Pi))
	//向下取整
	fmt.Printf("math.Floor(math.Pi): %v\n", math.Floor(math.Pi))
	//取绝对值
	fmt.Printf("math.Abs(-8): %v\n", math.Abs(-8))
	//开平方
	fmt.Printf("math.Sqrt(64): %v\n", math.Sqrt(64))
	fmt.Println("---------------------")
	//随机数,实际是个伪随机数,先设置伪随机种子
	rand.Seed(time.Now().UnixMicro())
	fmt.Printf("rand.Int(): %v\n", rand.Int())
	fmt.Printf("rand.Intn(100): %v\n", rand.Intn(3))   //产生100以内的随机数
	fmt.Printf("rand.Float32(): %v\n", rand.Float32()) //产生0-1之间的随机小数
}
