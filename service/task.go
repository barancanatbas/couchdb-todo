package service

import (
	"couchdb/model"
	repo "couchdb/repository"

	guid "github.com/google/uuid"
)

type ItaskService interface {
	GetAll() []model.Task
	Get(id string) model.Task
	Create(task model.Task) error
	Update(task model.Task) error
	Delete(id string) error
}

type TaskService struct {
	repo repo.ItaskRepository
}

func NewTaskService(repo repo.ItaskRepository) ItaskService {
	return &TaskService{repo: repo}
}

func (t TaskService) GetAll() []model.Task {
	return t.repo.GetAll()
}

func (t TaskService) Get(id string) model.Task {
	return t.repo.Get(id)
}

func (t TaskService) Create(task model.Task) error {
	id := guid.New().String()

	task.ID = id

	err := t.repo.Create(task)
	return err
}

func (t TaskService) Update(task model.Task) error {
	return t.repo.Update(task)
}

func (t TaskService) Delete(id string) error {
	return t.repo.Delete(id)
}
