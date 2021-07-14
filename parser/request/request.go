package request

import (
	"net/url"
)

type Analytics struct {
}

type Parser struct {
}

func (p *Parser) ParseRequest(in <-chan url.URL) <-chan Analytics {
	out := make(chan Analytics)
	// go func() {
	// 	defer close(out)
	// 	for t := range in {
	// 		fmt.Printf("task: %s, parse object %s from %s", t.ID, t.Key, t.InputBucket)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()
	return out
}
