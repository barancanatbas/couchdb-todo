package main

import (
	"couchdb/config"
	"couchdb/handler"
	"couchdb/repository"
	"couchdb/service"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	cluster := config.Connect()
	taskRepository := repository.NewTaskRepository(cluster)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	e.GET("/tasks", taskHandler.GetAll)
	e.GET("/task/:id", taskHandler.Get)
	e.POST("/task", taskHandler.Create)
	e.PUT("/task", taskHandler.Update)
	e.DELETE("/task/:id", taskHandler.Delete)

	e.Start(":8080")
}
