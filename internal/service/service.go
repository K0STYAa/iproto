package service

import (
	"github.com/K0STYAa/vk_iproto/internal/repository"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

type State interface {
	ChangeState(stateId uint8)
}

type ReadWrite interface {
	Read(ReqReadArgs models.ReqReadArgs) (models.RespReadArgs, error)
	Replace(ReqReplaceArgs models.ReqReplaceArgs) (models.RespReplaceArgs, error)
}

type Service struct {
	State
	ReadWrite
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		State:     NewStateService(repos.State),
		ReadWrite: NewReadWriteService(repos.ReadWrite),
	}
}
