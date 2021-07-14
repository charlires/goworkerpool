package s3

import (
	"fmt"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetObject(bucket, key string) ([]byte, error) {
	fmt.Println("getting file")
	time.Sleep(100 * time.Millisecond)
	return nil, nil
}

func (s *Service) PutObject(bucket, key string, body []byte) error {
	fmt.Println("putting file")
	time.Sleep(100 * time.Millisecond)
	return nil
}

// func (s *Service) GetObject(in <-chan workerpool.Task) <-chan workerpool.Task {
// 	out := make(chan workerpool.Task)
// 	go func() {
// 		defer close(out)
// 		for t := range in {
// 			fmt.Printf("task: %s, downloading object %s from %s", t.ID, t.Key, t.InputBucket)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()
// 	return out
// }

// func (s *Service) PutObject(in <-chan workerpool.Task) <-chan workerpool.Task {
// 	out := make(chan workerpool.Task)
// 	go func() {
// 		defer close(out)
// 		for t := range in {
// 			fmt.Printf("task: %s, uploading object %s to %s", t.ID, t.Key, t.OutputBucket)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()
// 	return out
// }
