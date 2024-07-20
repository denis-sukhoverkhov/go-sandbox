package sandbox

import (
	"log"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Duration time.Duration
}

type Worker struct {
	ID int
	wg *sync.WaitGroup
}

func (w Worker) ProcessTask(t Task) {
	log.Printf("Worker %d processing task %d\n", w.ID, t.ID)

	time.Sleep(time.Millisecond * 100)
}

func (w Worker) Start(taskQ <-chan Task) {
	defer w.wg.Done()
	for t := range taskQ {
		w.ProcessTask(t)
	}
}

type WorkerPool struct {
	WorkerCt  int
	TaskQueue chan Task
	Workers   []Worker
	wg        *sync.WaitGroup
}

func NewWorkerPool(workerCt int) *WorkerPool {
	pool := &WorkerPool{
		WorkerCt:  workerCt,
		TaskQueue: make(chan Task),
		wg:        &sync.WaitGroup{},
	}
	pool.initWorkers()
	return pool
}

func (r *WorkerPool) initWorkers() {
	r.Workers = make([]Worker, 0, r.WorkerCt)

	r.wg.Add(r.WorkerCt)
	for i := 0; i < r.WorkerCt; i++ {
		r.Workers = append(r.Workers, Worker{ID: i + 1, wg: r.wg})
		go r.Workers[len(r.Workers)-1].Start(r.TaskQueue)
	}
}

func (r *WorkerPool) Wait() {
	r.wg.Wait()
}
