package collection

import (
	"encoding/json"
	"fmt"
	"sync"

	"golang.org/x/exp/maps"
)

type CSet[T comparable] struct {
	set map[T]struct{}
	mu  sync.RWMutex
}

func (s *CSet[T]) Contains(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.set[item]
	return exists
}

func (s *CSet[T]) Insert(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set[item] = struct{}{}
}

func (s *CSet[T]) Delete(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.set, item)
}

func (s *CSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.set)
}

func (s *CSet[T]) UnmarshalJSON(data []byte) error {
	var temp []T
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	s.set = map[T]struct{}{}
	for _, val := range temp {
		s.Insert(val)
	}
	return nil
}

func (s CSet[T]) String() string {
	keys := maps.Keys(s.set)
	return fmt.Sprintf("%v", keys)
}
