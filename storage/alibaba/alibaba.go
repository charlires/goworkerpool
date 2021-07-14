package alibaba

import (
	"fmt"
	"time"

	"github.com/Joker666/goworkerpool/workerpool"
)

type Service struct {
}

func (s *Service) GetObject(in <-chan workerpool.Task) <-chan workerpool.Task {
	out := make(chan workerpool.Task)
	go func() {
		defer close(out)
		for t := range in {
			fmt.Printf("task: %s, downloading object %s from %s", t.ID, t.Key, t.InputBucket)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

func (s *Service) PutObject(in <-chan workerpool.Task) <-chan workerpool.Task {
	out := make(chan workerpool.Task)
	go func() {
		defer close(out)
		for t := range in {
			fmt.Printf("task: %s, uploading object %s to %s", t.ID, t.Key, t.OutputBucket)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}
