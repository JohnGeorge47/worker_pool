package worker

import (
	"fmt"
	"time"
)

type Worker struct {
	ID          int
	KillChan    chan struct{}
	Work        chan Job
	WorkerQueue chan chan Job
}

func NewWorker(id int, workerQueue chan chan Job) Worker {
	worker := Worker{
		ID:          id,
		KillChan:    make(chan struct{}),
		Work:        make(chan Job),
		WorkerQueue: workerQueue,
	}
	return worker
}

func (w *Worker) Start() {
	go func() {
		w.WorkerQueue <- w.Work
		select {
		case work := <-w.Work:
			job_details := fmt.Sprintf("The Job id is %s and value is %s", work.JobId, work.Value)
			time.Sleep(5*time.Second)
			worker_details := fmt.Sprintf("The worker id is %s", w.ID)
			fmt.Println(job_details, worker_details)
		case <-w.KillChan:
			fmt.Printf("worker %d stopping", w.ID)
			return
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.KillChan <- struct{}{}
	}()
}
