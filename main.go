package main

import (
	"gin-crud/controller"
	"gin-crud/services"
)

func main() {
	services := services.NewServices()
	router := controller.NewRouter(services)

	router.Run(":5000")
}
