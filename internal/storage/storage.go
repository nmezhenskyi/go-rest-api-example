// Package storage implements in-memory hash table storage.
package storage

import "sync"

type Storage struct {
	table map[string]interface{}
	mu    sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{table: make(map[string]interface{})}
}

func (s *Storage) FindById(id string) (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.table[id]

	return value, ok
}

func (s *Storage) FindAll() ([]interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.table) == 0 {
		return []interface{}{}, false
	}

	items := make([]interface{}, 0, len(s.table))
	for _, value := range s.table {
		items = append(items, value)
	}

	return items, true
}

func (s *Storage) Save(id string, item interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.table[id] = item
}

func (s *Storage) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.table, id)
}

func (s *Storage) SetEmpty() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.table = make(map[string]interface{})
}
