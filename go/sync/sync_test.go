package sync

import (
	"strconv"
	"sync"
	"testing"
)

type RWMutexMap struct {
	m  map[interface{}]interface{}
	mu sync.RWMutex
}

func (m *RWMutexMap) Delete(key interface{}) {
	_, ok := m.Load(key)
	if !ok {
		return
	}
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *RWMutexMap) Load(key interface{}) (value interface{}, ok bool) {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}

	m.mu.RLock()
	v, ok := m.m[key]
	m.mu.RUnlock()
	return v, ok
}

func (m *RWMutexMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	v, ok := m.Load(key)
	if ok {
		return v, true
	}
	m.Store(key, value)
	return value, false
}

func (m *RWMutexMap) Range(f func(key, value interface{}) bool) {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}

	m.mu.RLock()
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
	m.mu.RUnlock()
}

func (m *RWMutexMap) Store(key, value interface{}) {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}

	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

type Map interface {
	Delete(key interface{})
	Load(key interface{}) (value interface{}, ok bool)
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Range(f func(key, value interface{}) bool)
	Store(key, value interface{})
}

func benchmark_Map(m Map) {
	var wg sync.WaitGroup
	// Store
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				m.Store(strconv.Itoa(i), i)
			}
		}()
	}

	// Range
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				m.Range(func(k, v interface{}) bool {
					return true
				})
			}
		}()
	}
	wg.Wait()

	// Delete
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				m.Delete(strconv.Itoa(i))
			}
		}()
	}

	// LoadOrStore
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				m.LoadOrStore(strconv.Itoa(i), i)
			}
		}()
	}
}

func Benchmark_Map(b *testing.B) {
	b.Run("sync.Map", func(b *testing.B) {
		m := new(sync.Map)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchmark_Map(m)
		}
	})

	b.Run("sync.RWMutex", func(b *testing.B) {
		m := new(RWMutexMap)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchmark_Map(m)
		}
	})
}
