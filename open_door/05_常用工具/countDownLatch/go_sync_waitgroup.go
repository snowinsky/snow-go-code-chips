package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("start ....")
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)
	go task2(&waitGroup)
	go task4(&waitGroup)
	go task6(&waitGroup)
	waitGroup.Wait() //阻塞，直至这3个任务全部执行完毕
	log.Println("end ....")

}

func task2(ws *sync.WaitGroup) {
	defer ws.Done()
	time.Sleep(2 * time.Second)
}

func task4(ws *sync.WaitGroup) {
	defer ws.Done()
	time.Sleep(4 * time.Second)
}

func task6(ws *sync.WaitGroup) {
	defer ws.Done()
	time.Sleep(6 * time.Second)
}
