package test

import (
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/stretchr/testify/mock"
)

type ItemService struct {
	mock.Mock
}

func NewItemService(r repository.ItemRepository) *ItemService {
	return new(ItemService)
}

func (m *ItemService) ListItems() ([]models.Item, error) {
	args := m.Called()

	return args.Get(0).([]models.Item), args.Error(1)
}

func (m *ItemService) FindItem(id string) (*models.Item, bool, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Item), args.Get(1).(bool), args.Error(2)
}

func (m *ItemService) CreateItem(input models.CreateItemInput) (*models.Item, error) {
	args := m.Called(input)
	return args.Get(0).(*models.Item), args.Error(1)
}

func (m *ItemService) UpdateItem(id string, input models.UpdateItemInput) (*models.Item, error) {
	args := m.Called(id, input)
	return args.Get(0).(*models.Item), args.Error(1)
}

func (m *ItemService) DeleteItem(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ItemService) SetItemRepository(r repository.ItemRepository) {
	m.Called(r)
}

func (m *ItemService) GetItemRepository() repository.ItemRepository {
	args := m.Called()
	return args.Get(0).(repository.ItemRepository)
}
