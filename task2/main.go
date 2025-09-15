package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Hello! it is task2 begin")
	// 指针测试
	// i := 1
	// y := pointer_test(&i)
	// fmt.Println("调用后 pointer_test 返回：", *y)

	// 切片测试
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println("original slice:", nums)
	// slice_test(&nums)
	// fmt.Println("modify slice:", nums)

	//go 关键字启动两个协程
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go printOdds(&wg)
	// go printEvens(&wg)
	// wg.Wait()
	// fmt.Println("所有协程执行完毕")

	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	// 初始化调度器，设置最大并发数为3，总任务数为10
	// scheduler := NewScheduler(3, 10)
	// // 添加10个示例任务
	// for i := 1; i <= 10; i++ {
	// 	taskDuration := time.Duration((i%3)+1) * time.Second // 任务耗时1-3秒
	// 	task := Task{
	// 		ID:      i,
	// 		Handler: sampleTaskFunc(i, taskDuration),
	// 	}
	// 	scheduler.AddTask(task)
	// }
	// // 启动结果收集器
	// go scheduler.CollectResults()
	// // 启动调度器
	// scheduler.Start()
	// // 等待所有任务完成
	// time.Sleep(5 * time.Second) // 给任务时间完成
	// // 打印摘要
	// scheduler.PrintSummary()

	//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
	// rect := Rectangle{Height: 5, Width: 3}
	// circle := Circle{Radius: 4}
	// fmt.Println("Rectangle Area:", rect.Area())
	// fmt.Println("Rectangle Perimeter:", rect.Perimeter())
	// fmt.Println("Circle Area:", circle.Area())
	// fmt.Println("Circle Perimeter:", circle.Perimeter())

	//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
	// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	// emp := Employee{EmployeeId: 1, Person: Person{Name: "张三", Age: 30}}
	// emp.PrintInfo()
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go producer(ch, &wg)
	// go consumer(ch, &wg)
	// wg.Wait()
	// fmt.Println("所有数据已处理完成")

	// bufferSize := 100
	// totalItems := 100
	// ch := make(chan int, bufferSize)
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go consumer1(ch, &wg)
	// producer1(ch, totalItems)
	// wg.Wait()
	// fmt.Println("所有数据已处理完成")

	// wg := sync.WaitGroup{}
	// safeCounter := &safeCounter{}
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for j := 0; j < 1000; j++ {
	// 			safeCounter.inc()
	// 		}
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println("计数器结果：", safeCounter.getCount())
	//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var counter int64
	couter1 := int64(0)
	wg := sync.WaitGroup{}
	fmt.Println("couter1:", couter1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("最终计数值: %d\n", atomic.LoadInt64(&counter))
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func slice_test(nums *[]int) {
	fmt.Println("slice_test :进入参数为：", nums)
	slice_nums := *nums
	for i := 0; i < len(slice_nums); i++ {
		// slice_nums[i] = slice_nums[i] * 2
		slice_nums[i] *= 2
	}

}

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func pointer_test(i *int) *int {
	fmt.Println(*i)
	*i = *i + 10
	return i
}
