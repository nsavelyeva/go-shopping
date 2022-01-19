package services

// The services layer is responsible for the business logic of the application.
// The service layer will delegate reading and writing data to the repositories and external API clients,
// so that it can focus on the business logic.

import (
	"log"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/repository"
)

type ItemService interface {
	ListItems() ([]models.Item, error)
	FindItem(id string) (*models.Item, bool, error)
	CreateItem(input models.CreateItemInput) (*models.Item, error)
	UpdateItem(id string, input models.UpdateItemInput) (*models.Item, error)
	DeleteItem(id string) error
	SetItemRepository(r repository.ItemRepository)
}

type itemService struct {
    r *repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) *ItemService {
	if r == nil {
		log.Fatal("Failed to initialize item service, repository is nil")
		return nil
	}
	var s ItemService = &itemService{r: &r}
	return &s
}

func (s *itemService) SetItemRepository(r repository.ItemRepository) {
	s.r = &r
}

func (s *itemService) GetItemRepository() repository.ItemRepository {
	if s.r == nil {
		log.Fatal("Failed to get item repository, it is nil")
		return nil
	}

	return *s.r
}

func (s *itemService) ListItems() ([]models.Item, error) {
	items, err := s.GetItemRepository().ListItems()
	return items, err
}

func (s *itemService) FindItem(id string) (*models.Item, bool, error) {
	r := s.GetItemRepository()
	item, found, err := r.FindItem(id)
	return item, found, err
}

func (s *itemService) CreateItem(input models.CreateItemInput) (*models.Item, error) {
	// Assumed input is validated on upper (handlers) layer
	r := s.GetItemRepository()
	item, err := r.CreateItem(&input)
	return item, err
}

func (s *itemService) UpdateItem(id string, input models.UpdateItemInput) (*models.Item, error) {
	// Assumed input is validated on upper (handlers) layer
	r := s.GetItemRepository()
	item, err := r.UpdateItem(id, &input)
	return item, err
}

func (s *itemService) DeleteItem(id string) error {
	r := s.GetItemRepository()
	err := r.DeleteItem(id)
	return err
}
