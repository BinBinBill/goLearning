package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型
type Task struct {
	ID      int
	Handler func() error
}

// TaskResult 存储任务执行结果
type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Error     error
}

// Scheduler 任务调度器
type Scheduler struct {
	maxConcurrent  int             // 最大并发数
	taskQueue      chan Task       // 任务队列
	resultQueue    chan TaskResult // 结果队列
	wg             sync.WaitGroup  // 等待组
	totalTasks     int             // 总任务数
	completedTasks int             // 已完成任务数
	taskResults    []TaskResult    // 任务结果列表
	mutex          sync.Mutex      // 互斥锁
}

// NewScheduler 创建新的调度器
func NewScheduler(maxConcurrent, totalTasks int) *Scheduler {
	return &Scheduler{
		maxConcurrent:  maxConcurrent,
		taskQueue:      make(chan Task, totalTasks),
		resultQueue:    make(chan TaskResult, totalTasks),
		totalTasks:     totalTasks,
		completedTasks: 0,
		taskResults:    make([]TaskResult, 0, totalTasks),
	}
}

// AddTask 添加任务到队列
func (s *Scheduler) AddTask(task Task) {
	s.taskQueue <- task
}

// worker 工作goroutine，执行任务
func (s *Scheduler) worker(id int) {
	defer s.wg.Done()

	for task := range s.taskQueue {
		start := time.Now()

		// 执行任务
		err := task.Handler()

		end := time.Now()
		duration := end.Sub(start)

		// 发送结果到结果队列
		s.resultQueue <- TaskResult{
			TaskID:    task.ID,
			StartTime: start,
			EndTime:   end,
			Duration:  duration,
			Error:     err,
		}
	}
}

// Start 启动调度器
func (s *Scheduler) Start() {
	// 启动worker
	for i := 0; i < s.maxConcurrent; i++ {
		s.wg.Add(1)
		go s.worker(i)
	}

	// 等待所有任务完成并关闭任务队列
	go func() {
		s.wg.Wait()
		close(s.resultQueue)
	}()
}

// CollectResults 收集并处理结果
func (s *Scheduler) CollectResults() {
	for result := range s.resultQueue {
		s.mutex.Lock()
		s.taskResults = append(s.taskResults, result)
		s.completedTasks++

		fmt.Printf("任务%d完成: 开始时间=%s, 结束时间=%s, 耗时=%v, 错误=%v\n",
			result.TaskID, result.StartTime.Format("15:04:05.000"),
			result.EndTime.Format("15:04:05.000"), result.Duration, result.Error)

		// 实时显示完成进度
		fmt.Printf("进度: %d/%d (%.1f%%)\n",
			s.completedTasks, s.totalTasks,
			float64(s.completedTasks)/float64(s.totalTasks)*100)
		s.mutex.Unlock()
	}
}

// PrintSummary 打印执行摘要
func (s *Scheduler) PrintSummary() {
	var totalDuration time.Duration
	var successCount int

	for _, result := range s.taskResults {
		totalDuration += result.Duration
		if result.Error == nil {
			successCount++
		}
	}

	avgDuration := totalDuration / time.Duration(len(s.taskResults))

	fmt.Println("\n===== 任务执行摘要 =====")
	fmt.Printf("总任务数: %d\n", s.totalTasks)
	fmt.Printf("成功任务数: %d\n", successCount)
	fmt.Printf("失败任务数: %d\n", s.totalTasks-successCount)
	fmt.Printf("总执行时间: %v\n", totalDuration)
	fmt.Printf("平均任务耗时: %v\n", avgDuration)
}

// 示例任务函数
func sampleTaskFunc(id int, duration time.Duration) func() error {
	return func() error {
		fmt.Printf("任务%d开始执行，将持续%s...\n", id, duration)
		time.Sleep(duration)

		// 模拟偶尔失败
		if id%7 == 0 { // 让ID能被7整除的任务失败
			return fmt.Errorf("任务%d模拟失败", id)
		}

		fmt.Printf("任务%d完成\n", id)
		return nil
	}
}
