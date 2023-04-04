package repositories

import (
	"github.com/cvitorin/complero/ims/src/domain"
)

type ItemRepository interface {
	GetAll() ([]*domain.Item, error)
}

type DummyItemRepository struct {
}

func (r *DummyItemRepository) GetAll() ([]*domain.Item, error) {
	return []*domain.Item{
		domain.NewItem("Rock", 10, 12),
		domain.NewItem("Aged Cheese", 5, 10),
		domain.NewItem("Conjured Sword", 10, 5),
		domain.NewItem("Regular Sword", 2, 10),
	}, nil
}
