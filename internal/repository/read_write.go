package repository

import (
	"errors"

	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

type ReadWriteStorage struct {
	storage *internal.BaseStorage
}

func NewReadWriteStorage(storage *internal.BaseStorage) *ReadWriteStorage {
	return &ReadWriteStorage{storage: storage}
}

func (s *ReadWriteStorage) Read(ReqReadArgs models.ReqReadArgs) (models.RespReadArgs, error) {
	if s.storage.StorageState == models.Maintenance {
		return models.RespReadArgs{S: ""}, errors.New("Can't Read At Maintenance Mode")
	}
	if ReqReadArgs.Id < models.FirstDataId || ReqReadArgs.Id > models.StorageDataLen-1 {
		return models.RespReadArgs{S: ""}, errors.New("Invalid ID. Valid Value in range[0; 999]")
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	return models.RespReadArgs{S: s.storage.Data[ReqReadArgs.Id]}, nil
}

func (s *ReadWriteStorage) Replace(ReqReplaceArgs models.ReqReplaceArgs) (models.RespReplaceArgs, error) {
	if s.storage.StorageState == models.Maintenance {
		return models.RespReplaceArgs{}, errors.New("Can't Replace At Maintenance Mode")
	}
	if s.storage.StorageState == models.ReadOnly {
		return models.RespReplaceArgs{}, errors.New("Can't Replace At ReadOnly Mode")
	}
	if len(ReqReplaceArgs.S) > models.StringMaxLen {
		return models.RespReplaceArgs{}, errors.New("Incoming String Cannot Take Up More Than 256 Bytes")
	}
	if ReqReplaceArgs.Id < models.FirstDataId || ReqReplaceArgs.Id > models.StorageDataLen-1 {
		return models.RespReplaceArgs{}, errors.New("Invalid ID. Valid Value in range[0; 999]")
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	s.storage.Data[ReqReplaceArgs.Id] = ReqReplaceArgs.S
	return models.RespReplaceArgs{}, nil
}
