package main

import (
	"github.com/fvbock/endless"
	"github.com/nsavelyeva/go-shopping/routers"
	"time"
)

func main() {
	router_ := routers.Setup()
	time.Now().Sub(x)
	endless.ListenAndServe(":8080", router) // instead of router.Run()
	//if err != nil {
	//	return
	//}
}
