package crdtgo

import (
	"sync"
)

type CRDT struct {
	crdt map[interface{}]struct{}
	lock sync.RWMutex
}

func New(items ...interface{}) *CRDT {
	set := &CRDT{
		crdt: make(map[interface{}]struct{}),
	}

	for _, item := range items {
		set.Add(item)
	}

	return set
}

func (set *CRDT) Add(item interface{}) {
	set.lock.Lock()
	defer set.lock.Unlock()
	set.crdt[item] = struct{}{}
}

func (set *CRDT) Remove(item interface{}) {
	set.lock.Lock()
	defer set.lock.Unlock()
	delete(set.crdt, item)
}

func (set *CRDT) Size() int {
	set.lock.RLock()
	defer set.lock.RUnlock()
	return len(set.crdt)
}
