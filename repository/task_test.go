package repository

import (
	"couchdb/config"
	"couchdb/model"
	"testing"
)

func TestGet(t *testing.T) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	task := taskRepository.Get("4c98b983-edf4-4c8c-b014-096328f11c7b")
	if task.ID != "4c98b983-edf4-4c8c-b014-096328f11c7b" {
		t.Errorf("Expected 4c98b983-edf4-4c8c-b014-096328f11c7b, got %s", task.ID)
	}
}

func TestGetAll(t *testing.T) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	tasks := taskRepository.GetAll()
	if len(tasks) < 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestCreateDocument(t *testing.T) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	task := model.Task{
		ID:   "sdlkfsjdf",
		Name: "Test",
		Done: false,
	}

	err := taskRepository.Create(task)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestUpdateDocument(t *testing.T) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	task := model.Task{
		ID:   "4c98b983-edf4-4c8c-b014-096328f11c7b",
		Name: "Test",
		Done: false,
	}

	err := taskRepository.Update(task)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestDeleteDocument(t *testing.T) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	err := taskRepository.Delete("4c98b983-edf4-4c8c-b014-096328f11c7b")
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}
