package storage

import (
	"errors"

	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/pkg/iproto"
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

func (s *ReadWriteStorage) Read(reqReadArgs iproto.ReqReadArgs) (iproto.RespReadArgs, error) {
	if s.storage.StorageState == internal.Maintenance {
		return iproto.RespReadArgs{S: ""}, ErrReadMaintenanceMode
	}

	if reqReadArgs.ID < internal.FirstDataID || reqReadArgs.ID > internal.StorageDataLen-1 {
		return iproto.RespReadArgs{S: ""}, ErrInvalidID
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	return iproto.RespReadArgs{S: s.storage.Data[reqReadArgs.ID]}, nil
}

func (s *ReadWriteStorage) Replace(reqReplaceArgs iproto.ReqReplaceArgs) (iproto.RespReplaceArgs, error) {
	if s.storage.StorageState == internal.Maintenance {
		return iproto.RespReplaceArgs{}, ErrReplaceMaintenanceMode
	}

	if s.storage.StorageState == internal.ReadOnly {
		return iproto.RespReplaceArgs{}, ErrReplaceReadOnlyMode
	}

	if len(reqReplaceArgs.S) > internal.StringMaxLen {
		return iproto.RespReplaceArgs{}, ErrBigString
	}

	if reqReplaceArgs.ID < internal.FirstDataID || reqReplaceArgs.ID > internal.StorageDataLen-1 {
		return iproto.RespReplaceArgs{}, ErrInvalidID
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	s.storage.Data[reqReplaceArgs.ID] = reqReplaceArgs.S

	return iproto.RespReplaceArgs{}, nil
}
