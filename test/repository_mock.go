package test

import (
    "github.com/nsavelyeva/go-shopping/models"
    "github.com/stretchr/testify/mock"
)

type ItemRepository struct {
    mock.Mock
}

func NewItemRepository(driverName string, connectionString string) *ItemRepository {
    return new(ItemRepository)
}

func (m *ItemRepository) ListItems() ([]models.Item, error) {
    args := m.Called()

    return args.Get(0).([]models.Item), args.Error(1)
}

func (m *ItemRepository) FindItem(id string) (*models.Item, bool, error) {
    args := m.Called(id)
    return args.Get(0).(*models.Item), args.Get(1).(bool), args.Error(2)
}

func (m *ItemRepository) CreateItem(input *models.CreateItemInput) (*models.Item, error) {
    args := m.Called(input)
    return args.Get(0).(*models.Item), args.Error(1)
}

func (m *ItemRepository) UpdateItem(id string, input *models.UpdateItemInput) (*models.Item, error) {
    args := m.Called(id, input)
    return args.Get(0).(*models.Item), args.Error(1)
}

func (m *ItemRepository) DeleteItem(id string) error {
    args := m.Called(id)
    return args.Error(0)
}
