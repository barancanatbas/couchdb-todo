package repository

import (
	"couchdb/model"

	"github.com/couchbase/gocb"
)

type ItaskRepository interface {
	GetAll() []model.Task
	Get(id string) model.Task
	Create(task model.Task) error
	Update(task model.Task) error
	Delete(id string) error
}

type TaskRepository struct {
	bucket *gocb.Bucket
}

func NewTaskRepository(cluster *gocb.Cluster) ItaskRepository {
	bucket, _ := cluster.OpenBucket("tasks", "")
	return &TaskRepository{bucket: bucket}
}

func (t TaskRepository) GetAll() []model.Task {
	var tasks []model.Task
	var task model.Task

	query := gocb.NewN1qlQuery("SELECT ID,Name,Done FROM tasks")
	rows, err := t.bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		panic(err)
	}

	for rows.Next(&task) {
		tasks = append(tasks, task)
		task = model.Task{}
	}

	return tasks
}

func (t TaskRepository) Get(id string) model.Task {
	var task model.Task

	query := gocb.NewN1qlQuery("SELECT ID,Name,Done FROM tasks WHERE ID = $1")
	rows, err := t.bucket.ExecuteN1qlQuery(query, []interface{}{id})
	if err != nil {
		panic(err)
	}

	for rows.Next(&task) {
		return task
	}

	return task
}

func (t TaskRepository) Create(task model.Task) error {
	_, err := t.bucket.Insert(task.ID, task, 0)

	if err != nil {
		return err
	}
	return nil
}

func (t TaskRepository) Update(task model.Task) error {
	_, err := t.bucket.Replace(task.ID, task, 0, 0)
	if err != nil {
		return err
	}

	return nil
}

func (t TaskRepository) Delete(id string) error {
	_, err := t.bucket.Remove(id, 0)
	if err != nil {
		return err
	}

	return nil
}
