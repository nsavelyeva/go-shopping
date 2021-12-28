package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/database"
	"github.com/nsavelyeva/go-shopping/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIEnv struct {
	DB *gorm.DB
}

// GET /items
// List all items
func (a *APIEnv) ListItems(c *gin.Context) {
	items, err := database.ListItems(a.DB)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// GET /items/:id
// Find an item
func (a *APIEnv) FindItem(c *gin.Context) {
	item, found, err := database.FindItem(a.DB, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}
	if found {
		c.JSON(http.StatusOK, gin.H{"item": item})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"item": nil})
}

// POST /items
// Create new item
func (a *APIEnv) CreateItem(c *gin.Context) {
	// Validate input
	var input models.CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create item
	item, err := database.CreateItem(a.DB, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"item": item})
}

// PATCH /items/:id
// Update an item
func (a *APIEnv) UpdateItem(c *gin.Context) {
	// Validate input
	var input models.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := database.UpdateItem(a.DB, c.Param("id"), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"item": item})
}

// DELETE /items/:id
// Delete an item
func (a *APIEnv) DeleteItem(c *gin.Context) {
	if err := database.DeleteItem(a.DB, c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
