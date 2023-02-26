package repository

import (
	"github.com/K0STYAa/vk_iproto/internal"
)

type StateStorage struct {
	storage *internal.BaseStorage
}

func NewStateStorage(storage *internal.BaseStorage) *StateStorage {
	return &StateStorage{storage: storage}
}

func (s *StateStorage) ChangeState(stateId uint8) {
	s.storage.Mutex.Lock()
	defer s.storage.Mutex.Unlock()

	s.storage.StorageState = stateId
}
