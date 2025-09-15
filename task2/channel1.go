package main

import (
	"fmt"
	"sync"
)

func producer1(ch chan<- int, total int) {
	defer close(ch) // 生产者完成发送后关闭通道
	for i := 1; i <= total; i++ {
		ch <- i // 向缓冲通道发送数据
		fmt.Printf("生产者发送: %d\n", i)
	}
}
func consumer1(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()       // 消费者完成时通知 WaitGroup
	for num := range ch { // 通过 range 自动感知通道关闭
		fmt.Printf("消费者接收: %d\n", num)
	}
}
