package storage

import (
	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/pkg/iproto"
)

type State interface {
	ChangeState(stateID uint8)
}

type ReadWrite interface {
	Read(ReqReadArgs iproto.ReqReadArgs) (iproto.RespReadArgs, error)
	Replace(ReqReplaceArgs iproto.ReqReplaceArgs) (iproto.RespReplaceArgs, error)
}

type Repository struct {
	State
	ReadWrite
}

func NewRepository(storage *internal.BaseStorage) *Repository {
	return &Repository{
		State:     NewStateStorage(storage),
		ReadWrite: NewReadWriteStorage(storage),
	}
}
