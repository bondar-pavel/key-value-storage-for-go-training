package storage

import (
	"sync"
)

type Storage struct {
	Map map[string]string
	Mutex *sync.Mutex
}

func (s *Storage) Init() {
	s.Map = make(map[string]string)
	s.Mutex = &sync.Mutex{}
}

func (s *Storage) Set(key, value string) {
	s.Mutex.Lock()
	s.Map[key] = value
	s.Mutex.Unlock()
}

func (s *Storage) Get(key string) (string, bool) {
	s.Mutex.Lock()
	value, ok := s.Map[key]
	s.Mutex.Unlock()
	return value, ok
}

func (s *Storage) Delete(key string) (deleted bool) {
	s.Mutex.Lock()
	_, exists := s.Map[key]
	if exists == true {
		delete(s.Map, key)
		deleted = true
	}
	s.Mutex.Unlock()
	return
}
