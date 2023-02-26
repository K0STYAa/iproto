package service

import "github.com/K0STYAa/vk_iproto/internal/repository"

type StateService struct {
	repo repository.State
}

func NewStateService(repo repository.State) *StateService {
	return &StateService{repo: repo}
}

func (s *StateService) ChangeState(stateID uint8) {
	s.repo.ChangeState(stateID)
}
