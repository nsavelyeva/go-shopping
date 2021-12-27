package main

import (
	"github.com/nsavelyeva/go-shopping/controllers"
	"github.com/nsavelyeva/go-shopping/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/items", controllers.ListItems)
	r.GET("/items/:id", controllers.FindItem)
	r.POST("/items", controllers.CreateItem)
	r.PATCH("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)

	// Run the server
	r.Run()
}
