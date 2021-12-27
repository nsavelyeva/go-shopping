package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/models"
)

type CreateItemInput struct {
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

type UpdateItemInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Sold  bool    `json:"sold"`
}

// GET /items
// List all items
func ListItems(c *gin.Context) {
	var items []models.Item
	models.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{"items": items})
}

// GET /items/:id
// Find an item
func FindItem(c *gin.Context) {
	// Get model if exist
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

// POST /items
// Create new item
func CreateItem(c *gin.Context) {
	// Validate input
	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create item
	item := models.Item{Name: input.Name, Price: input.Price, Sold: false}
	models.DB.Create(&item)

	c.JSON(http.StatusOK, gin.H{"item": item})
}

// PATCH /items/:id
// Update an item
func UpdateItem(c *gin.Context) {
	// Get model if exist
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}

	// Validate input
	var input UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&item).Updates(input)

	c.JSON(http.StatusOK, gin.H{"item": item})
}

// DELETE /items/:id
// Delete an item
func DeleteItem(c *gin.Context) {
	// Get model if exist
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}

	models.DB.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
