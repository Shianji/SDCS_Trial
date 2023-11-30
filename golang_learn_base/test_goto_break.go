package main

import "fmt"

func main() {
	// go语言中的goto语句和C语言一样
MARK:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue MARK
			}
			fmt.Printf("i=%v,j=%v  ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue
			}
			fmt.Printf("i=%v,j=%v  ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				break
			}
			fmt.Printf("i=%v,j=%v  ", i, j)
		}
		fmt.Println()
	}

}
