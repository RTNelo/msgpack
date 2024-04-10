package common

import (
	"sync"
	"sync/atomic"
)

type Map[K comparable, V any] struct {
	m atomic.Pointer[map[K]V]
	l sync.Mutex
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	v, ok := (*m.m.Load())[key]
	return v, ok
}

func (m *Map[K, V]) Store(key K, value V) {
	m.l.Lock()
	orig := *m.m.Load()
	newMap := make(map[K]V, len(orig)+1)
	for k, v := range orig {
		newMap[k] = v
	}
	newMap[key] = value
	m.m.Store(&newMap)
	m.l.Unlock()
}

func (m *Map[K, V]) Delete(key K) {
	m.l.Lock()
	orig := *m.m.Load()
	newMap := make(map[K]V, len(orig)-1)
	for k, v := range orig {
		if k != key {
			newMap[k] = v
		}
	}
	m.m.Store(&newMap)
	m.l.Unlock()
}

func NewMap[K comparable, V any]() *Map[K, V] {
	m := make(map[K]V)
	ret := &Map[K, V]{
		m: atomic.Pointer[map[K]V]{},
		l: sync.Mutex{},
	}
	ret.m.Store(&m)
	return ret
}
