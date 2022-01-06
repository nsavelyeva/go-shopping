package handlers
// The handler layer is responsible for parsing a request,
// calling out the relevant service and then returning a response to the caller.

import (
	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/nsavelyeva/go-shopping/services"
	"net/http"
)

type Provider struct {
	s *services.ItemService
}

func NewProvider(s services.ItemService, r repository.ItemRepository) *Provider {
	if s == nil {
		s = *services.NewItemService(r)
	} else {
		s.SetItemRepository(r)
	}
	var p Provider
	p.SetItemService(s)
	return &p
}

func (p *Provider) SetItemService(s services.ItemService) {
	p.s = &s
}

func (p *Provider) GetItemService(r repository.ItemRepository) services.ItemService {
	if p.s != nil {
		return *p.s
	}

	s := services.NewItemService(r)

	return *s
}

// GET /items - List all items
func (p *Provider) ListItems(c *gin.Context) {
	s := p.GetItemService(nil)
	items, err := s.ListItems()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

// GET /items/:id - Find an item
func (p *Provider) FindItem(c *gin.Context) {
	s := p.GetItemService(nil)
	item, found, err := s.FindItem(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}
	if found {
		c.JSON(http.StatusOK, gin.H{"data": item})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"data": nil})
}

// POST /items - Create a new item
func (p *Provider) CreateItem(c *gin.Context) {
	// Validate input
	var input models.CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an item
	s := p.GetItemService(nil)
	item, err := s.CreateItem(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

// PATCH /items/:id - Update an item
func (p *Provider) UpdateItem(c *gin.Context) {
	// Validate input
	var input models.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := p.GetItemService(nil)
	item, err := s.UpdateItem(c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

// DELETE /items/:id - Delete an item
func (p *Provider) DeleteItem(c *gin.Context) {
	s := p.GetItemService(nil)
	if err := s.DeleteItem(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
