package workerpool

import (
	"fmt"
	"time"
)

type InputStorageProvider interface {
	GetObject(bucket, key string) ([]byte, error)
	// GetObject(in <-chan Task) <-chan Task
}

type OutputStorageProvider interface {
	PutObject(bucket, key string, body []byte) error
	// PutObject(in <-chan Task) <-chan Task
}

type FileParserProvider interface {
	ParseObject(body []byte) ([]byte, error)
	// ParseObject(in <-chan Task) <-chan Task
}

type AuditProvider interface {
	Audit(*Task) error
}

// Task encapsulates a work item that should go in a work
type Task struct {
	ID           string // uuid
	Key          string // location
	InputBucket  string // s3/alibaba bucket
	OutputBucket string // s3/alibaba bucket
	CreatedAt    time.Time
	StartedAt    time.Time
	FinishedAt   time.Time
	isp          InputStorageProvider
	osp          OutputStorageProvider
	fp           FileParserProvider
	ap           AuditProvider
	Err          error
}

// NewTask initializes a new task based on a given work function.
func NewTask(id string, isp InputStorageProvider, osp OutputStorageProvider, fp FileParserProvider, ap AuditProvider) *Task {
	return &Task{ID: id, CreatedAt: time.Now(), isp: isp, osp: osp, fp: fp, ap: ap}
}

func (t *Task) Do(workerID int) {
	t.StartedAt = time.Now()
	defer func() {
		t.FinishedAt = time.Now()
		err := t.ap.Audit(t)
		if err != nil {
			fmt.Printf("unable to log tast %s, err:%s\n", t.ID, err)
		}
	}()

	fmt.Printf("Worker %d starting task %s\n", workerID, t.ID)

	var err error
	var b []byte
	b, err = t.isp.GetObject(t.InputBucket, t.Key)
	if err != nil {
		t.Err = fmt.Errorf("downloading object %w", err)
		return
	}
	var b2 []byte
	b2, err = t.fp.ParseObject(b)
	if err != nil {
		t.Err = fmt.Errorf("parsing object %w", err)
		return
	}
	err = t.osp.PutObject(t.InputBucket, t.Key, b2)
	if err != nil {
		t.Err = fmt.Errorf("uploading object %w", err)
		return
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finishing task %s took %d\n", workerID, t.ID, time.Since(t.StartedAt))
}
