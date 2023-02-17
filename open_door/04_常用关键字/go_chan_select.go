package main

import (
	"fmt"
	"strconv"
	"time"
)

//select 下设多个case，一个default，如果有一个case准备好了，就执行，如果都没准备好就走default
//没有default时，如果所有case都没准备好，就等待；如果都准备好了，就随机执行-个
//select是一个选择执行器，只执行一个；可以设想的场景是，有多个consumer同时消费一个channel中的数据，哪个consumer消费到了，就走哪个

func main() {
	var channel1 = make(chan string, 2)
	for i := 0; i < 10; i++ {
		select {
		case s := <-channel1:
			fmt.Println(s)
		case channel1 <- "a":
			fmt.Println("produce a complete")
		}
	}

	var channel2 = make(chan string, 20)
	producer1(channel2)
	for {
		select {
		case s := <-channel2:
			fmt.Println("consumer1 " + s)
		case s := <-channel2:
			fmt.Println("consumer2 " + s)
		case s := <-channel2:
			fmt.Println("consumer3 " + s)
		case <-time.After(10 * time.Second): //如果都消费不到消息，且超过10秒，可以进入超时处理
			fmt.Println("consumer timeout")
		}
	}
}

// 生产者
func producer1(channel1 chan<- string) {
	for i := 0; i < 10; i++ {
		channel1 <- "P0" + strconv.Itoa(i)
	}
}

func consumer1(channel1 <-chan string) {
	s := <-channel1
	fmt.Println("consumer1 " + s)
}

func consumer2(channel1 <-chan string) {
	s := <-channel1
	fmt.Println("consumer2 " + s)
}

func consumer3(channel1 <-chan string) {
	s := <-channel1
	fmt.Println("consumer3 " + s)
}
