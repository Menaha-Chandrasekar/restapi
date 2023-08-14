package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID   int
	Name string
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d is working on task %d: %s\n", id, task.ID, task.Name)
	}
}

func main() {
	numWorkers := 6
	numTasks := 10

	var wg sync.WaitGroup
	tasks := make(chan Task, numTasks)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i, Name: fmt.Sprintf("Task %d", i)}
	}

	close(tasks)
	wg.Wait()
}
