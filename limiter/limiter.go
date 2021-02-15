package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	sync.Mutex
	counter int
	Limit   int
}

func New(maxLimit int) *Limiter {
	l := &Limiter{}
	l.Limit = maxLimit
	go func() {
		for {
			time.Sleep(time.Second)
			l.Lock()
			l.counter = 0
			l.Unlock()
		}
	}()
	return l
}

func (l *Limiter) Check() bool {
	defer l.Unlock()
	l.Lock()
	l.counter++
	if l.counter > l.Limit {
		return false
	}
	return true
}
