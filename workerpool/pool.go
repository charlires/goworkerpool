package workerpool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Pool is the worker pool
type Pool struct {
	Tasks   []*Task
	Workers []*Worker

	concurrency   int //number of workers
	waitingTime   time.Duration
	Collector     chan *Task
	runBackground chan bool
	wg            sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewPool(
	tasks []*Task, concurrency, tasksBuffer int, waitingTime time.Duration,
) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		Collector:   make(chan *Task, tasksBuffer),
		waitingTime: waitingTime,
	}
}

// AddTask adds a task to the pool
func (p *Pool) AddTask(task *Task) {
	p.Collector <- task
}

// RunBackground runs the pool in background
func (p *Pool) Run(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		for {
			fmt.Print("⌛ Waiting for tasks to come in ...\n")
			time.Sleep(p.waitingTime)
		}
	}()

	wg.Add(p.concurrency)
	for i := 1; i <= p.concurrency; i++ {
		worker := NewWorker(p.Collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.Start(ctx, wg)
	}

	for i := range p.Tasks {
		p.Collector <- p.Tasks[i]
	}

}
