package service

import (
	"ecommerce/model"
	"ecommerce/repository"
	"errors"
)

type ItemService struct {
	repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

// AddItem adds a new item to the inventory.
func (service *ItemService) AddItem(item *model.Item) error {
	// Validate item data
	if item.Title == "" || item.Cost <= 0 || item.Quantity < 0 {
		return errors.New("invalid item data")
	}

	// Insert item into the database
	return service.repo.AddItem(item)
}

// GetItemByID retrieves an item by its ID.
func (service *ItemService) GetItemByID(id int) (*model.Item, error) {
	item, err := service.repo.GetItemByID(id)
	if err != nil {
		return nil, errors.New("item not found")
	}
	return item, nil
}

// UpdateItem updates an item's details.
func (service *ItemService) UpdateItem(item *model.Item) error {
	// Validate item data
	if item.Title == "" || item.Cost <= 0 || item.Quantity < 0 {
		return errors.New("invalid item data")
	}

	// Update item in the database
	return service.repo.UpdateItem(item)
}

// DeleteItem deletes an item from the inventory.
func (service *ItemService) DeleteItem(id int) error {
	return service.repo.DeleteItem(id)
}

// GetAllItems retrieves all items with pagination.
func (service *ItemService) GetAllItems(page, limit int) ([]model.Item, error) {
	return service.repo.GetAllItems(page, limit)
}
