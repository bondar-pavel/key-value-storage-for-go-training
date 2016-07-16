package storage

type Storage struct {
	Map map[string]string
}

func (s *Storage) Init() {
	s.Map = make(map[string]string)
}

func (s *Storage) Set(key, value string) {
	s.Map[key] = value
}

func (s *Storage) Get(key string) (string, bool) {
	value, ok := s.Map[key]
	return value, ok
}

func (s *Storage) Delete(key string) bool {
	_, exists := s.Map[key]
	if exists == true {
		delete(s.Map, key)
		return true
	}
	return false
}
