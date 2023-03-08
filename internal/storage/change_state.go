package storage

type StateStorage struct {
	storage *BaseStorage
}

func NewStateStorage(storage *BaseStorage) *StateStorage {
	return &StateStorage{storage: storage}
}

func (s *StateStorage) ChangeState(stateID uint8) {
	s.storage.Mutex.Lock()
	defer s.storage.Mutex.Unlock()

	s.storage.StorageState = stateID
}
