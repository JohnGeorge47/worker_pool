package dispatcher

import (
	"fmt"
	"github.com/JohnGeorge47/sendx/pkg/worker"
	"sync"
)

var WorkerQueue chan chan worker.Job

type Dispatcher struct {
	Mu sync.Mutex
	Size int
	WorkerQueue chan chan worker.Job
	Workers []*worker.Worker
	JobQueue chan worker.Job
}

func StartDispatcher(nworkers int)*Dispatcher{
	WorkerQueue = make(chan chan worker.Job, 100)
	d:=&Dispatcher{
		Size:        nworkers,
		WorkerQueue: WorkerQueue,
		Workers:make([]*worker.Worker,nworkers,100),
		JobQueue:make(chan worker.Job),
	}
	for i := 0; i < nworkers; i++ {
		worker := worker.NewWorker(i+1, WorkerQueue)
		fmt.Println("starting worker %d", worker.ID)
		worker.Start()
		d.Workers[i]=&worker
	}
	fmt.Println(d.Workers)
	go func() {
		for {
			select {
			case work := <-d.JobQueue:
				fmt.Println("got a work request")
				go func() {
					worker,ok := <-d.WorkerQueue
					fmt.Println(ok)
					worker <- work
				}()
			}
		}
	}()
	return d
}

func (d Dispatcher)Resize(worker_count int){
  d.Mu.Lock()
  defer d.Mu.Unlock()
  if d.Size<worker_count{
  	fmt.Println("here1")
  	fmt.Println(d.Size,worker_count)
	  for d.Size<worker_count {
		  d.Size++
		  worker:=worker.NewWorker(d.Size+1,d.WorkerQueue)
		  fmt.Println("Starting worker:",worker.ID)
		  worker.Start()
		  d.Workers=append(d.Workers, &worker)
		  fmt.Println(d.Workers)
	  }
  }else if d.Size>worker_count {
  	   fmt.Println(d.Size)
  	    toremove:=d.Size-worker_count
		arr_toremove,current_arr:=d.Workers[0:toremove],d.Workers[toremove:len(d.Workers)]
		fmt.Println(arr_toremove)
	  for i, w := range arr_toremove {
		  fmt.Println(i,w.ID)
		  w.Stop()
	  }
	  d.Size=worker_count
	  d.Workers=current_arr
	  fmt.Println(d.Workers)
	  fmt.Println(d.Size)
  }
	return
}