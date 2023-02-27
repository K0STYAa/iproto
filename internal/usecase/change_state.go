package usecase

import "github.com/K0STYAa/vk_iproto/internal/storage"

type StateUsecase struct {
	repo storage.State
}

func NewStateUsecase(repo storage.State) *StateUsecase {
	return &StateUsecase{repo: repo}
}

func (s *StateUsecase) ChangeState(stateID uint8) {
	s.repo.ChangeState(stateID)
}
