package inmem

import "sync"

//Package inmem holds in-memory key/value data protected by a mutex
//GGOD : sync.Mutex is a named, unexported field, it stays private

type InMem struct {
	//named + unexported — Lock/Unlock are NOT visible outside
	mu sync.Mutex
	m  map[string]int
}

func New() *InMem {
	return &InMem{
		m: make(map[string]int),
	}
}

func (i *InMem) Get(key string) (int, bool) {
	i.mu.Lock()
	v, ok := i.m[key]
	i.mu.Unlock()
	return v, ok
}

func (i *InMem) Set(key string, val int) {
	i.mu.Lock()
	i.m[key] = val
	i.mu.Unlock()
}
