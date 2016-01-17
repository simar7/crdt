package crdtgo

import (
	"fmt"
	"strings"
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

func (set *CRDT) Exists(item interface{}) bool {
	set.lock.RLock()
	defer set.lock.Unlock()
	_, state := set.crdt[item]
	return state
}

func (set *CRDT) Get(elem interface{}) interface{} {
	set.lock.RLock()
	defer set.lock.RUnlock()

	for iter, item := range set.List() {
		if elem == item {
			return iter
		}
	}
	return nil
}

func (set *CRDT) Size() int {
	set.lock.RLock()
	defer set.lock.RUnlock()
	return len(set.crdt)
}

func (set *CRDT) List() []interface{} {
	set.lock.RLock()
	defer set.lock.RUnlock()
	list := make([]interface{}, 0, len(set.crdt))
	for item := range set.crdt {
		list = append(list, item)
	}

	return list
}

func (set *CRDT) String() string {
	str := make([]string, 0, len(set.List()))
	for _, item := range set.List() {
		str = append(str, fmt.Sprintf("%v", item))
	}

	return fmt.Sprintf("[%s]", strings.Join(str, ", "))
}
