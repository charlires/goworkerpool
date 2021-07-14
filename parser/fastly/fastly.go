package fastly

import (
	"fmt"
	"net/url"
	"time"

	"github.com/Joker666/goworkerpool/parser/request"
)

type RequestParser interface {
	ParseRequest(in <-chan url.URL) <-chan request.Analytics
}

type Parser struct {
	requestParser RequestParser
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParseObject([]byte) ([]byte, error) {
	fmt.Println("parsing file")
	time.Sleep(100 * time.Millisecond)
	return nil, nil
}

// func (p *Parser) ParseObject(in <-chan workerpool.Task) <-chan workerpool.Task {
// 	out := make(chan workerpool.Task)
// 	go func() {
// 		defer close(out)
// 		for t := range in {
// 			fmt.Printf("task: %s, parse object %s from %s", t.ID, t.Key, t.InputBucket)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()
// 	return out
// }
