package services

import (
	"github.com/cvitorin/complero/ims/src/domain"
	"github.com/cvitorin/complero/ims/src/repositories"
)

type ItemService interface {
	GetAllItems() ([]*domain.Item, error)
}

type DefaultItemService struct {
	ItemRepository repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) *DefaultItemService {
	return &DefaultItemService{ItemRepository: repo}
}

func (s *DefaultItemService) GetAllItems() ([]*domain.Item, error) {
	return s.ItemRepository.GetAll()
}
