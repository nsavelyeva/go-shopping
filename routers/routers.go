package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/handlers"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/nsavelyeva/go-shopping/services"
)

func Setup() *gin.Engine {
	router := gin.Default()
	var r = *repository.NewItemRepository("sqlite3", "items.db")
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
