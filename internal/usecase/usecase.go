package usecase

import (
	"github.com/K0STYAa/vk_iproto/internal/storage"
	"github.com/K0STYAa/vk_iproto/pkg/iproto"
)

type State interface {
	ChangeState(stateID uint8)
}

type ReadWrite interface {
	Read(ReqReadArgs iproto.ReqReadArgs) (iproto.RespReadArgs, error)
	Replace(ReqReplaceArgs iproto.ReqReplaceArgs) (iproto.RespReplaceArgs, error)
}

type Usecase struct {
	State
	ReadWrite
}

func NewUsecase(repos *storage.Repository) *Usecase {
	return &Usecase{
		State:     NewStateUsecase(repos.State),
		ReadWrite: NewReadWriteUsecase(repos.ReadWrite),
	}
}
