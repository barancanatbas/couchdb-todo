package handler

import (
	"couchdb/model"
	"couchdb/service"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service service.ItaskService
}

func NewTaskHandler(service service.ItaskService) TaskHandler {
	return TaskHandler{service: service}
}

func (h TaskHandler) GetAll(c echo.Context) error {
	tasks := h.service.GetAll()
	return c.JSON(200, tasks)
}

func (h TaskHandler) Get(c echo.Context) error {
	id := c.Param("id")
	task := h.service.Get(id)
	return c.JSON(200, task)
}

func (h TaskHandler) Create(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return err
	}
	if err := h.service.Create(*task); err != nil {
		return err
	}
	return c.JSON(200, task)
}

func (h TaskHandler) Update(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return err
	}
	if err := h.service.Update(*task); err != nil {
		return err
	}
	return c.JSON(200, task)
}

func (h TaskHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		return err
	}
	return c.JSON(200, "success !")
}
