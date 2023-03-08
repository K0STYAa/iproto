package storage

import "github.com/K0STYAa/iproto/pkg/iproto"

type State interface {
	ChangeState(stateID uint8)
}

type ReadWrite interface {
	Read(ReqReadArgs iproto.ReqReadArgs) (iproto.RespReadArgs, error)
	Replace(ReqReplaceArgs iproto.ReqReplaceArgs) (iproto.RespReplaceArgs, error)
}

type Storage struct {
	State
	ReadWrite
}

func NewStorage(storage *BaseStorage) *Storage {
	return &Storage{
		State:     NewStateStorage(storage),
		ReadWrite: NewReadWriteStorage(storage),
	}
}
