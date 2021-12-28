package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"github.com/nsavelyeva/go-shopping/controllers"
	"github.com/nsavelyeva/go-shopping/models"
)

func main() {
	router := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/items", controllers.ListItems)
	router.GET("/items/:id", controllers.FindItem)
	router.POST("/items", controllers.CreateItem)
	router.PATCH("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)

	// Run the server
	endless.ListenAndServe(":8080", router)   // instead of router.Run()
}
