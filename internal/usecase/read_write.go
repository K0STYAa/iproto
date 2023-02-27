package usecase

import (
	"fmt"

	"github.com/K0STYAa/vk_iproto/internal/storage"
	"github.com/K0STYAa/vk_iproto/pkg/iproto"
)

type ReadWriteUsecase struct {
	repo storage.ReadWrite
}

func NewReadWriteUsecase(repo storage.ReadWrite) *ReadWriteUsecase {
	return &ReadWriteUsecase{repo: repo}
}

const errTemplate = "%w"

func (s *ReadWriteUsecase) Read(reqReadArgs iproto.ReqReadArgs) (iproto.RespReadArgs, error) {
	resp, err := s.repo.Read(reqReadArgs)

	if err != nil {
		return resp, fmt.Errorf(errTemplate, err)
	}

	return resp, nil
}

func (s *ReadWriteUsecase) Replace(reqReplaceArgs iproto.ReqReplaceArgs) (iproto.RespReplaceArgs, error) {
	resp, err := s.repo.Replace(reqReplaceArgs)

	if err != nil {
		return resp, fmt.Errorf(errTemplate, err)
	}

	return resp, nil
}
