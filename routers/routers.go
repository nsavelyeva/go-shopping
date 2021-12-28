package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/database"
	"github.com/nsavelyeva/go-shopping/handlers"
)

func Setup() *gin.Engine {
	router := gin.Default()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}
	// Routes
	router.GET("/items", api.ListItems)
	router.GET("/items/:id", api.FindItem)
	router.POST("/items", api.CreateItem)
	router.PATCH("/items/:id", api.UpdateItem)
	router.DELETE("/items/:id", api.DeleteItem)

	return router
}
