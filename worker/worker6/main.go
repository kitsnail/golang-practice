package main

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id  int
	Err error
	f   func() error
}

func (task *Task) Do() error {
	return task.f()
}

type WorkerPool struct {
	PoolSize   int
	taskSize   int
	taskChan   chan Task
	resultChan chan Task
	Results    func() []Task
}

func NewWorkerPool(tasks []Task, size int) *WorkerPool {
	tasksChan := make(chan Task, len(tasks))
	resultsChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)
	pool := &WorkerPool{
		PoolSize:   size,
		taskSize:   len(tasks),
		taskChan:   tasksChan,
		resultChan: resultsChan,
	}
	pool.Results = pool.results
	return pool
}

func (p *WorkerPool) results() []Task {
	tasks := make([]Task, p.taskSize)
	for i := 0; i < p.taskSize; i++ {
		tasks[i] = <-p.resultChan
	}
	return tasks
}

func (p *WorkerPool) Start() {
	for i := 0; i < p.PoolSize; i++ {
		go p.worker()
	}
}

func (p *WorkerPool) worker() {
	for task := range p.taskChan {
		task.Err = task.Do()
		p.resultChan <- task
	}
}

func main() {
	t := time.Now()

	tasks := []Task{
		{Id: 0, f: func() error { time.Sleep(2 * time.Second); fmt.Println(0); return nil }},
		{Id: 1, f: func() error { time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},
		{Id: 2, f: func() error { fmt.Println(2); return errors.New("error") }},
	}
	pool := NewWorkerPool(tasks, 2)
	pool.Start()

	tasks = pool.Results()
	fmt.Printf("all tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
	for _, task := range tasks {
		fmt.Printf("result of task %d is %v\n", task.Id, task.Err)
	}
}
