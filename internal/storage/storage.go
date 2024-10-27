package storage

func NewStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]UserData),
	}
}

func (s *MemoryStorage) Get(userID string) (UserData, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, ok := s.data[userID]
	return data, ok
}

func (s *MemoryStorage) Set(userID string, data UserData) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[userID] = data
}

func (s *MemoryStorage) Delete(userID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, userID)
}
