package main

import (
	"sync"

	"github.com/nice-pink/nicelog/log"
)

type Test struct {
	Prefix string
}

func (t Test) Test(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Warn("test message", "prefix", t.Prefix)
}
