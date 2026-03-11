package inmem

import "sync"

//Embedded - promotes Lock() and Unlock() publicly
type InMem struct {
	sync.Mutex
	m map[string]int
}

func New() *InMem {
	return &InMem{
		m: make(map[string]int),
	}
}

func (i *InMem) Get(key string) (int, bool) {
	i.Lock()
	v, ok := i.m[key]
	i.Unlock()
	return v, ok
}

func (i *InMem) Set(key string, val int) {
	i.Lock()
	i.m[key] = val
	i.Unlock()
}
