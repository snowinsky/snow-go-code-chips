package main

import (
	"fmt"
	"strconv"
	"time"
)

//对于消费者来说，面临的问题如下：
//1. 什么时候该停止空等？
//2. 什么叫消费完了？

//解决方案
//1. 等待接受数据的同时，不停的查看数据源是否关闭（这种方式效率太低，属于）

// 生产者
func producer(channel1 chan<- string) {
	for i := 0; i < 10; i++ {
		channel1 <- "P0" + strconv.Itoa(i)
	}
}

// 消费者 使用range遍历channel中的元素，直到没了为止
func consumer(channel1 <-chan string) {
	for s := range channel1 {
		fmt.Println("consume the " + s)
	}
}

func main() {
	var channel1 = make(chan string, 2)
	go consumer(channel1)
	go producer(channel1)

	time.Sleep(10 * time.Second)

}
