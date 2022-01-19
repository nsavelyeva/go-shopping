package routers

// TODO: move usage of gorm

import (
	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/handlers"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/nsavelyeva/go-shopping/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup is a function to initiate gin and define routes, it is used in test
func Setup() *gin.Engine {
	router := gin.Default()
	var r = *repository.NewItemRepository(sqlite.Open("items.db"), &gorm.Config{})
	var s = services.NewItemService(r)
	var h = handlers.NewItemHandler(*s)
	// Routes
	router.GET("/items", h.ListItems)
	router.GET("/items/:id", h.FindItem)
	router.POST("/items", h.CreateItem)
	router.PATCH("/items/:id", h.UpdateItem)
	router.DELETE("/items/:id", h.DeleteItem)

	return router
}
