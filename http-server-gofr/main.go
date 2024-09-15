package main

import "gofr.dev/pkg/gofr"

func eventhandler(*gofr.Context) (interface{}, error) {
	return "Welcome to gofr server", nil
}

func main() {
	app := gofr.New() // App is of the type GoFR app

	app.GET("/event", eventhandler)
	app.POST("/event", eventhandler)
	app.Run()
}
