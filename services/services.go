package services
// The services layer is responsible for the business logic of the application.
// The service layer will delegate reading and writing data to the repositories and external API clients,
// so that it can focus on the business logic.

import (
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/repository"
)

type ItemService interface {
	ListItems() ([]models.Item, error)
	FindItem(id string) (models.Item, bool, error)
	CreateItem(input models.CreateItemInput) (models.Item, error)
	UpdateItem(id string, input models.UpdateItemInput) (models.Item, error)
	DeleteItem(id string) error
	SetItemRepository(r repository.ItemRepository)
}

type itemService struct {
    r *repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) *ItemService {
	if r == nil {
		r = *repository.NewItemRepository()
	}
	var p itemService
	p.SetItemRepository(r)
	var s ItemService = &p

	return &s
}

func (s *itemService) SetItemRepository(r repository.ItemRepository) {
	s.r = &r
}

func (s *itemService) GetItemRepository() repository.ItemRepository {
	if s.r != nil {
		return *s.r
	}

	r := repository.NewItemRepository()

	return *r
}

// GET /items - List all items
func (s *itemService) ListItems() ([]models.Item, error) {
	items, err := s.GetItemRepository().ListItems()
	return items, err
}

// GET /items/:id - Find an item
func (s *itemService) FindItem(id string) (models.Item, bool, error) {
	r := s.GetItemRepository()
	item, found, err := r.FindItem(id)
	return item, found, err
}

// POST /items - Create a new item
func (s *itemService) CreateItem(input models.CreateItemInput) (models.Item, error) {
	// Assumed input is validated on upper (handlers) layer
	r := s.GetItemRepository()
	item, err := r.CreateItem(&input)
	return *item, err
}

// PATCH /items/:id - Update an item
func (s *itemService) UpdateItem(id string, input models.UpdateItemInput) (models.Item, error) {
	// Assumed input is validated on upper (handlers) layer
	r := s.GetItemRepository()
	item, err := r.UpdateItem(id, &input)
	return *item, err
}

// DELETE /items/:id - Delete an item
func (s *itemService) DeleteItem(id string) error {
	r := s.GetItemRepository()
	err := r.DeleteItem(id)
	return err
}
