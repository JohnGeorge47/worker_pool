package worker_v2

import (
	"fmt"
	"time"
)

type Job struct {
	Id int
	Value interface{}
}

type WorkRequest struct {
	JobChan chan <-Job
}

type Worker struct {
	WorkerId int
	ReqChan chan <- WorkRequest
	CloseChan chan struct{}
}

func NewWorker(id int,reqchan  chan <- WorkRequest)*Worker{
	w:=Worker{
		WorkerId: id ,
		ReqChan:   reqchan,
		CloseChan: make(chan struct{}),
	}
	go w.run()
	return &w
}

func (w Worker)run(){
	job_chan:=make(chan Job)
	for  {
		select {
		case w.ReqChan<-WorkRequest{JobChan:job_chan}:
			time.Sleep(5 * time.Second)
			current_job:=<-job_chan
			fmt.Println(current_job)
		case <-w.CloseChan:
			return
		}
	}

}

func (w *Worker)Stop(){
	close(w.CloseChan)
}