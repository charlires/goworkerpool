package workerpool

import (
	"context"
	"fmt"
	"sync"
)

// Worker handles all the work
type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan bool
}

// NewWorker returns new instance of worker
func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: channel,
		quit:     make(chan bool),
	}
}

// StartBackground starts the worker in background waiting
func (wr *Worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting worker %d\n", wr.ID)

	for {
		select {
		case task := <-wr.taskChan:
			task.Do(wr.ID)
		case <-ctx.Done():
			fmt.Printf("Worker %d received cancellation signal!\n", wr.ID)
			return
		}
	}
}
