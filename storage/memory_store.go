package storage

import (
	"errors"
	"fmt"
	"sync"
)

type Record map[string]interface{}

type Store struct {
	sync.RWMutex
	Data  map[string]map[string]Record // contentType -> id -> record
	Count map[string]int               // contentType -> auto-increment id
}

var MemoryStore = Store{
	Data:  make(map[string]map[string]Record),
	Count: make(map[string]int),
}

func (s *Store) Create(contentType string, record Record) string {
	s.Lock()
	defer s.Unlock()

	s.Count[contentType]++
	id := s.Count[contentType]
	idStr := fmt.Sprintf("%d", id)

	if s.Data[contentType] == nil {
		s.Data[contentType] = make(map[string]Record)
	}
	s.Data[contentType][idStr] = record
	return idStr
}

func (s *Store) Get(contentType, id string) (Record, error) {
	s.RLock()
	defer s.RUnlock()

	if rec, ok := s.Data[contentType][id]; ok {
		return rec, nil
	}
	return nil, errors.New("record not found")
}

func (s *Store) GetAll(contentType string) []Record {
	s.RLock()
	defer s.RUnlock()

	records := []Record{}
	for _, rec := range s.Data[contentType] {
		records = append(records, rec)
	}
	return records
}

func (s *Store) Update(contentType, id string, updated Record) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Data[contentType][id]; ok {
		s.Data[contentType][id] = updated
		return nil
	}
	return errors.New("record not found")
}

func (s *Store) Delete(contentType, id string) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Data[contentType][id]; ok {
		delete(s.Data[contentType], id)
		return nil
	}
	return errors.New("record not found")
}
