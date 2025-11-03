package storage

import "sync"

// In memory storage
// Thread-safe in-memory store (map + mutex)
type MemoryStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urls: make(map[string]string),
	}
}

// Save short ID -> original URL
func (m *MemoryStore) Save(short, original string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.urls[short] = original
}

// Get original URL by short ID
func (m *MemoryStore) Get(short string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.urls[short]
	return val, ok
}
