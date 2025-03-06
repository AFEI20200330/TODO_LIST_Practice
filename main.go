package main

import (
	"TODO_LIST_Practice/app"
	config2 "TODO_LIST_Practice/config"
)

func main() {
	config := config2.NewConfig()

	app := &app.App{}

	app.InitializeRoutes(config)
	app.Run(":8081")
}
