package repository

import (
	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

type State interface {
	ChangeState(stateId uint8)
}

type ReadWrite interface {
	Read(ReqReadArgs models.ReqReadArgs) (models.RespReadArgs, error)
	Replace(ReqReplaceArgs models.ReqReplaceArgs) (models.RespReplaceArgs, error)
}

type Repository struct{
	State
	ReadWrite
}

func NewRepository(storage *internal.BaseStorage) *Repository {
	return &Repository{
		State: NewStateStorage(storage),
		ReadWrite: NewReadWriteStorage(storage),
	}
}