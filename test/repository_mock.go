package test

import (
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/stretchr/testify/mock"
)

// ItemRepository is a struct to be used as a mock repository object in tests
type ItemRepository struct {
	mock.Mock
}

// NewItemRepository is a mock constructor for ItemRepository struct
func NewItemRepository(driverName string, connectionString string) *ItemRepository {
	return new(ItemRepository)
}

// ListItems is a mock method for ItemRepository struct
func (m *ItemRepository) ListItems() ([]models.Item, error) {
	args := m.Called()
	return args.Get(0).([]models.Item), args.Error(1)
}

// FindItem is a mock method for ItemRepository struct
func (m *ItemRepository) FindItem(id int) (*models.Item, bool, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Item), args.Get(1).(bool), args.Error(2)
}

// CreateItem is a mock method for ItemRepository struct
func (m *ItemRepository) CreateItem(input *models.CreateItemInput) (*models.Item, error) {
	args := m.Called(input)
	return args.Get(0).(*models.Item), args.Error(1)
}

// UpdateItem is a mock method for ItemRepository struct
func (m *ItemRepository) UpdateItem(id int, input *models.UpdateItemInput) (*models.Item, error) {
	args := m.Called(id, input)
	return args.Get(0).(*models.Item), args.Error(1)
}

// DeleteItem is a mock method for ItemRepository struct
func (m *ItemRepository) DeleteItem(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
