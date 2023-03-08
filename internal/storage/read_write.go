package storage

import (
	"errors"

	"github.com/K0STYAa/iproto/internal"
	"github.com/K0STYAa/iproto/pkg/iproto"
)

type ReadWriteStorage struct {
	storage *BaseStorage
}

func NewReadWriteStorage(storage *BaseStorage) *ReadWriteStorage {
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
	if s.storage.StorageState == internal.StorageStateMaintenance {
		return iproto.RespReadArgs{S: ""}, ErrReadMaintenanceMode
	}

	if reqReadArgs.ID < firstDataID || reqReadArgs.ID > storageDataLen-1 {
		return iproto.RespReadArgs{S: ""}, ErrInvalidID
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	return iproto.RespReadArgs{S: s.storage.Data[reqReadArgs.ID]}, nil
}

func (s *ReadWriteStorage) Replace(reqReplaceArgs iproto.ReqReplaceArgs) (iproto.RespReplaceArgs, error) {
	if s.storage.StorageState == internal.StorageStateMaintenance {
		return iproto.RespReplaceArgs{}, ErrReplaceMaintenanceMode
	}

	if s.storage.StorageState == internal.StorageStateReadOnly {
		return iproto.RespReplaceArgs{}, ErrReplaceReadOnlyMode
	}

	if len(reqReplaceArgs.S) > stringMaxLen {
		return iproto.RespReplaceArgs{}, ErrBigString
	}

	if reqReplaceArgs.ID < firstDataID || reqReplaceArgs.ID > storageDataLen-1 {
		return iproto.RespReplaceArgs{}, ErrInvalidID
	}

	s.storage.Mutex.RLock()
	defer s.storage.Mutex.RUnlock()

	s.storage.Data[reqReplaceArgs.ID] = reqReplaceArgs.S

	return iproto.RespReplaceArgs{}, nil
}
