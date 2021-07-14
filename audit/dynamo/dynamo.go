package dynamo

import (
	"fmt"
	"time"

	"github.com/Joker666/goworkerpool/workerpool"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Audit(t *workerpool.Task) error {
	fmt.Printf("logging task %s\n", t.ID)
	time.Sleep(100 * time.Millisecond)
	return nil
}
