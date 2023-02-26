package service

import (
	"fmt"

	"github.com/K0STYAa/vk_iproto/internal/repository"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

type ReadWriteService struct {
	repo repository.ReadWrite
}

func NewReadWriteService(repo repository.ReadWrite) *ReadWriteService {
	return &ReadWriteService{repo: repo}
}

const errTemplate = "%w"

func (s *ReadWriteService) Read(reqReadArgs models.ReqReadArgs) (models.RespReadArgs, error) {
	resp, err := s.repo.Read(reqReadArgs)

	if err != nil {
		return resp, fmt.Errorf(errTemplate, err)
	}

	return resp, nil
}

func (s *ReadWriteService) Replace(reqReplaceArgs models.ReqReplaceArgs) (models.RespReplaceArgs, error) {
	resp, err := s.repo.Replace(reqReplaceArgs)

	if err != nil {
		return resp, fmt.Errorf(errTemplate, err)
	}

	return resp, nil
}
