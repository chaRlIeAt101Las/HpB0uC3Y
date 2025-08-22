// 代码生成时间: 2025-08-22 13:05:38
package main

import (
    "fmt"
    "time"
# 优化算法效率
    "github.com/robfig/cron/v3"
    "log"
)

// Job is a function that performs a task
type Job func()
# 扩展功能模块

// Scheduler is a type that schedules and runs jobs
# NOTE: 重要实现细节
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler creates a new scheduler with a cron scheduler
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(cron.WithSeconds()),
    }
}
# NOTE: 重要实现细节

// AddJob adds a job to the scheduler to run at the specified schedule
func (s *Scheduler) AddJob(sched string, job Job) error {
# FIXME: 处理边界情况
    _, err := s.cron.AddFunc(sched, job)
    if err != nil {
        return fmt.Errorf("failed to add job to scheduler: %w", err)
    }
    return nil
}
# 扩展功能模块

// Start starts the scheduler, it will block until the scheduler is stopped
func (s *Scheduler) Start() error {
    s.cron.Start()
    return nil
}

// Stop stops the scheduler
func (s *Scheduler) Stop() error {
# 改进用户体验
    s.cron.Stop()
    return nil
}

// JobExample is an example job that prints the current time
func JobExample() {
    fmt.Println("Job is running at: ", time.Now().Format(time.RFC1123))
}
# FIXME: 处理边界情况

func main() {
    scheduler := NewScheduler()
# NOTE: 重要实现细节
    // Add jobs to the scheduler
    if err := scheduler.AddJob("*/5 * * * * *", JobExample); err != nil {
        log.Fatalf("Failed to add job: %s", err)
    }

    // Start the scheduler
    if err := scheduler.Start(); err != nil {
        log.Fatalf("Failed to start scheduler: %s", err)
    }
    
    // The scheduler will run indefinitely, you would normally have a condition to stop it
    // For demonstration purposes, we will run it for 1 minute and then stop it
    time.Sleep(60 * time.Second)
    scheduler.Stop()
# FIXME: 处理边界情况
    fmt.Println("Scheduler stopped")
}
