package workerpool

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

// Worker represents a worker that can process tasks.
type Worker struct {
	// Channel to receive tasks.
	tasks <-chan string
	// WaitGroup to signal when the worker is done.
	wg *sync.WaitGroup
	// Channel to write results
	out chan string
}

func getHex(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// NewWorker creates a new worker.
func NewWorker(tasks <-chan string, wg *sync.WaitGroup, out chan string) *Worker {
	return &Worker{
		tasks: tasks,
		wg:    wg,
		out:   out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	// Ваш код
	for val := range w.tasks {
		w.out <- getHex(val)
	}
	w.wg.Done()
}
