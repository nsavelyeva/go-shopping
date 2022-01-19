package test

import (
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/stretchr/testify/mock"
)

// ItemService is a struct to be used as a mock service object in tests
type ItemService struct {
	mock.Mock
}

// NewItemService is a mock constructor for ItemService struct
func NewItemService(r repository.ItemRepository) *ItemService {
	return new(ItemService)
}

// ListItems is a mock method for ItemService struct
func (m *ItemService) ListItems() ([]models.Item, error) {
	args := m.Called()

	return args.Get(0).([]models.Item), args.Error(1)
}

// FindItem is a mock method for ItemService struct
func (m *ItemService) FindItem(id string) (*models.Item, bool, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Item), args.Get(1).(bool), args.Error(2)
}

// CreateItem is a mock method for ItemService struct
func (m *ItemService) CreateItem(input models.CreateItemInput) (*models.Item, error) {
	args := m.Called(input)
	return args.Get(0).(*models.Item), args.Error(1)
}

// UpdateItem is a mock method for ItemService struct
func (m *ItemService) UpdateItem(id string, input models.UpdateItemInput) (*models.Item, error) {
	args := m.Called(id, input)
	return args.Get(0).(*models.Item), args.Error(1)
}

// DeleteItem is a mock method for ItemService struct
func (m *ItemService) DeleteItem(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
