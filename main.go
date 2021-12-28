package main

import (
	"github.com/fvbock/endless"
	"github.com/nsavelyeva/go-shopping/database"
	"github.com/nsavelyeva/go-shopping/routers"
)


func main() {
	database.Setup()
	router := routers.Setup()

	endless.ListenAndServe(":8080", router)   // instead of router.Run()
}
