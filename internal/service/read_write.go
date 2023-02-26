package service

import (
	"github.com/K0STYAa/vk_iproto/internal/repository"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

type ReadWriteService struct {
	repo repository.ReadWrite
}

func NewReadWriteService(repo repository.ReadWrite) *ReadWriteService {
	return &ReadWriteService{repo: repo}
}

func (s *ReadWriteService) Read(ReqReadArgs models.ReqReadArgs) (models.RespReadArgs, error) {
	resp, err := s.repo.Read(ReqReadArgs)

	return resp, err
}

func (s *ReadWriteService) Replace(ReqReplaceArgs models.ReqReplaceArgs) (models.RespReplaceArgs, error) {
	resp, err := s.repo.Replace(ReqReplaceArgs)

	return resp, err
}
