package repository

import (
	"couchdb/config"
	"couchdb/model"
	"strconv"
	"testing"
)

// 439	   2412833 ns/op	   29101 B/op	     187 allocs/op
func BenchmarkGetAll(b *testing.B) {

	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		taskRepository.GetAll()
	}
}

//  402	   2521415 ns/op	   24550 B/op	     136 allocs/op
func BenchmarkGet(b *testing.B) {

	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		taskRepository.Get("4c98b983-edf4-4c8c-b014-096328f11c7b")
	}
}

// 2766	    365277 ns/op	    4680 B/op	      54 allocs/op
func BenchmarkCreateDocument(b *testing.B) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		task := model.Task{
			ID:   strconv.Itoa(i),
			Name: "Test",
			Done: false,
		}
		taskRepository.Create(task)
	}
}

// 9985	    115106 ns/op	    2699 B/op	      45 allocs/op
func BenchmarkDelete(b *testing.B) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		taskRepository.Delete("4c98b983-edf4-4c8c-b014-096328f11c7b")
	}
}

// 4813	    263622 ns/op	    4103 B/op	      55 allocs/op
func BenchmarkUpdate(b *testing.B) {
	cluster := config.Connect()
	taskRepository := NewTaskRepository(cluster)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		task := model.Task{
			ID:   "4c98b983-edf4-4c8c-b014-096328f11c7b",
			Name: "Test",
			Done: false,
		}
		taskRepository.Update(task)
	}
}
