package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/handlers"
)

func Setup() *gin.Engine {
	router := gin.Default()

	var h = handlers.NewProvider(nil, nil)
	// Routes
	router.GET("/items", h.ListItems)
	router.GET("/items/:id", h.FindItem)
	router.POST("/items", h.CreateItem)
	router.PATCH("/items/:id", h.UpdateItem)
	router.DELETE("/items/:id", h.DeleteItem)

	return router
}
