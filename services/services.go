package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/nsavelyeva/go-shopping/models"
	"net/http"
)

type ItemService interface {
	ListItems()
	FindItem()
	CreateItem()
	UpdateItem()
	DeleteItem()
}

type itemService struct {
	c *gin.Context
    r *repository.ItemRepository
	db *gorm.DB
}

func NewItemService() *ItemService {
	r := repository.NewItemRepository()

	var s ItemService = &itemService{r: r}
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

// GET /items
// List all items
func (s *itemService) ListItems() {
	items, err := s.r.ListItems()
	if err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.c.JSON(http.StatusOK, gin.H{"data": items})
}

// GET /items/:id
// Find an item
func (s *itemService) FindItem() {
	item, found, err := s.r.FindItem(s.db, s.c.Param("id"))
	if err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}
	if found {
		s.c.JSON(http.StatusOK, gin.H{"data": item})
		return
	}
	s.c.JSON(http.StatusNoContent, gin.H{"data": nil})
}

// POST /items
// Create new item
func (s *itemService) CreateItem() {
	// Validate input
	var input models.CreateItemInput
	if err := s.c.ShouldBindJSON(&input); err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create item
	item, err := s.r.CreateItem(s.db, &input)
	if err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.c.JSON(http.StatusOK, gin.H{"data": item})
}

// PATCH /items/:id
// Update an item
func (s *itemService) UpdateItem() {
	// Validate input
	var input models.UpdateItemInput
	if err := s.c.ShouldBindJSON(&input); err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := s.r.UpdateItem(s.db, s.c.Param("id"), &input)
	if err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.c.JSON(http.StatusOK, gin.H{"data": item})
}

// DELETE /items/:id
// Delete an item
func (s *itemService) DeleteItem() {
	if err := s.r.DeleteItem(s.db, s.c.Param("id")); err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}
	s.c.JSON(http.StatusOK, gin.H{"deleted": true})
}

