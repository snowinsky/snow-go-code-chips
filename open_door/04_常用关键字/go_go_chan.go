package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//直接执行这个函数，会阻塞一段时间
	log.Println("start....")
	threadSleep(5)
	log.Println("end....")
	//go就是运行一个新的线程，让这个函数异步执行
	log.Println("start....")
	go threadSleep(10)
	log.Println("end....")

	// 声明不带缓冲的通道
	channel1 := make(chan string)
	//像下面这样写会报运行时错误，这是因为创建的channel是没有缓冲的，这时写入channel的数据是直接给到读出的变量的
	//写和读都是同步阻塞动作，所以会一直阻塞直到能写入，但是因为channel根本没容量根本写不进去。所以直接报deadlock了
	/*channel1 <- "a" //fatal error: all goroutines are asleep - deadlock!
	var a = <-channel1
	fmt.Println(a)*/

	//如果初始化的channel是容量为零的，就必须先给这个channel配上消费者，然后再执行生产动作
	go func() {
		var a = <-channel1
		fmt.Println(a)
	}()
	channel1 <- "a"

	// 声明带10个缓冲的通道
	//这个是带缓冲区的，可以直接存入缓冲区中，所以不会报错，但是如果缓冲区存满了，消费者却没有消费掉，还是会报错的
	bufferChannel1 := make(chan string, 1)
	bufferChannel1 <- "b"
	//bufferChannel1 <- "b1" //fatal error: all goroutines are asleep - deadlock!
	var b = <-bufferChannel1
	fmt.Println(b)

	// 声明只读通道
	readOnlyChannel := make(<-chan string, 1)
	fmt.Println(readOnlyChannel)
	//readOnlyChannel <- "c" 编译报错，只读的不能写入
	//var c = <-readOnlyChannel //fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(c)

	// 声明只写通道
	writeOnlyChannel := make(chan<- string, 1)
	writeOnlyChannel <- "d"
	//var d = <-writeOnlyChannel 编译报错，只写的不能读出
	//fmt.Println(d)
}

// 这是一个会花很长时间执行的函数
func threadSleep(a int) {
	time.Sleep(time.Duration(a) * time.Second)
}
