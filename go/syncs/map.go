package syncs

import (
	"sync"
)

// Map implements a sync Map.
type Map struct {
	mu sync.RWMutex
	m  map[interface{}]interface{}
}

// Get ...
func (m *Map) Get(k interface{}) (interface{}, bool) {
	m.mu.RLock()
	if m.m == nil {
		return nil, false
	}
	v, ok := m.m[k]
	m.mu.RUnlock()
	return v, ok
}

// Set ...
func (m *Map) Set(k, v interface{}) {
	m.mu.Lock()
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}
	m.m[k] = v
	m.mu.Unlock()
}
