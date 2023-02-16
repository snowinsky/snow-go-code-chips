package main

import (
	"fmt"
	"strconv"
)

func main() {
	//go语言里只有for一种循环方式，没有while和do while
	for i := 0; i < 10; i++ {
		fmt.Println("for i=" + strconv.Itoa(i))
	}
	//while(express)的玩法
	j := 10
	for j > 7 {
		j--
		fmt.Println("for j=" + strconv.Itoa(j))
	}
	//while(true)的玩法
	for {
		fmt.Println("for true")
		break
	}
}
