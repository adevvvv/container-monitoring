package service

import (
	"container-monitoring/backend/model"
	"container-monitoring/backend/repository"
)

type PingService interface {
	GetAll() ([]model.PingStatus, error)
	GetByID(id int) (*model.PingStatus, error)
	Create(status *model.PingStatus) error
	Update(status *model.PingStatus) error
	Delete(id int) error
}

type PingServiceImpl struct {
	repo repository.PingRepository
}

func NewPingService(repo repository.PingRepository) *PingServiceImpl {
	return &PingServiceImpl{repo: repo}
}

func (s *PingServiceImpl) GetAll() ([]model.PingStatus, error) {
	return s.repo.GetAll()
}

func (s *PingServiceImpl) GetByID(id int) (*model.PingStatus, error) {
	return s.repo.GetByID(id)
}

func (s *PingServiceImpl) Create(status *model.PingStatus) error {
	return s.repo.Create(status)
}

func (s *PingServiceImpl) Update(status *model.PingStatus) error {
	return s.repo.Update(status)
}

func (s *PingServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}
