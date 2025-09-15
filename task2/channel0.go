package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i // 发送数据到通道
		fmt.Printf("生产者发送: %d\n", i)
	}
	close(ch) // 发送完成后关闭通道
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch { // 通过range持续接收直到通道关闭
		fmt.Printf("消费者接收: %d\n", num)
	}
}
