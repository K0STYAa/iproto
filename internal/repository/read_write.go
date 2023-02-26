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

var (
	ErrReadMaintenanceMode    = errors.New("can't read at maintenance mode")
	ErrInvalidID              = errors.New("invalid ID. Valid value in range[0; 999]")
	ErrReplaceMaintenanceMode = errors.New("can't replace at maintenance mode")
	ErrReplaceReadOnlyMode    = errors.New("can't replace at readOnly mode")
	ErrBigString              = errors.New("incoming string cannot take up more than 256 bytes")
)

func (s *ReadWriteStorage) Read(reqReadArgs models.ReqReadArgs) (models.RespReadArgs, error) {
	if s.storage.StorageState == models.Maintenance {
		return models.RespReadArgs{S: ""}, ErrReadMaintenanceMode
	}

	if reqReadArgs.ID < models.FirstDataID || reqReadArgs.ID > models.StorageDataLen-1 {
		return models.RespReadArgs{S: ""}, ErrInvalidID
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	return models.RespReadArgs{S: s.storage.Data[reqReadArgs.ID]}, nil
}

func (s *ReadWriteStorage) Replace(reqReplaceArgs models.ReqReplaceArgs) (models.RespReplaceArgs, error) {
	if s.storage.StorageState == models.Maintenance {
		return models.RespReplaceArgs{}, ErrReplaceMaintenanceMode
	}

	if s.storage.StorageState == models.ReadOnly {
		return models.RespReplaceArgs{}, ErrReplaceReadOnlyMode
	}

	if len(reqReplaceArgs.S) > models.StringMaxLen {
		return models.RespReplaceArgs{}, ErrBigString
	}

	if reqReplaceArgs.ID < models.FirstDataID || reqReplaceArgs.ID > models.StorageDataLen-1 {
		return models.RespReplaceArgs{}, ErrInvalidID
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	s.storage.Data[reqReplaceArgs.ID] = reqReplaceArgs.S

	return models.RespReplaceArgs{}, nil
}
