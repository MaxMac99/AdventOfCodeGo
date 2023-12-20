package shared

import (
	"sync"
)

type Task func() int

type WorkPool interface {
	AddTask(Task)
	CompleteTasks()
	Sum() int
}

type workPool struct {
	tasks   chan Task
	results chan int
	sum     chan int
	wg      *sync.WaitGroup
}

func (w workPool) CompleteTasks() {
	close(w.tasks)
}

func worker(tasks <-chan Task, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result := task()
		results <- result
	}
}

func collectResults(results chan int, response chan int) {
	total := 0
	for result := range results {
		total += result
	}
	response <- total
}

func (w workPool) AddTask(task Task) {
	w.tasks <- task
}

func (w workPool) Sum() int {
	w.wg.Wait()
	close(w.results)

	result := <-w.sum
	close(w.sum)
	return result
}

func StartNewWorkPool(numWorkers int) WorkPool {
	w := &workPool{
		tasks:   make(chan Task),
		results: make(chan int),
		sum:     make(chan int),
		wg:      new(sync.WaitGroup),
	}

	for i := 0; i < numWorkers; i++ {
		w.wg.Add(1)
		go worker(w.tasks, w.results, w.wg)
	}

	go collectResults(w.results, w.sum)
	return w
}
