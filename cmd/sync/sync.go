package sync

import "sync"

type Counter struct {
	mu         sync.Mutex
	countValue int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (count *Counter) Inc() {

	count.mu.Lock()
	defer count.mu.Unlock()
	count.countValue++
}

func (count *Counter) Value() int {
	return count.countValue
}
