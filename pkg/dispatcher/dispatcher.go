package dispatcher

import (
	"fmt"
	"github.com/JohnGeorge47/sendx/pkg/worker"
)

var WorkerQueue chan chan worker.Job

type Dispatcher struct {
	Size int
	WorkerQueue chan chan worker.Job
}

func StartDispatcher(nworkers int, WorkQueue chan worker.Job) {
	WorkerQueue = make(chan chan worker.Job, nworkers+1)
	for i := 0; i < nworkers; i++ {
		worker := worker.NewWorker(i+1, WorkerQueue)
		fmt.Println("starting worker %d", worker.ID)
		worker.Start()
	}
	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("got a work request")
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}

func (d Dispatcher)Resize(){

}