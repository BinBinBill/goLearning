package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printOdds(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("协程1（奇数）: %d\n", i)
	}
}

func printEvens(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("协程2（偶数）: %d\n", i)
	}
}

//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
