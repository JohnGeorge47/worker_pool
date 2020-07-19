package worker_pool_v2

import (
	"github.com/JohnGeorge47/sendx/pkg/worker_v2"
	"sync"
)

type Pool struct {
	Workers []*worker_v2.Worker
	ReqChan chan worker_v2.WorkRequest
	Mu sync.Mutex
}

func NewPool(n int)*Pool{
	p:=&Pool{
		ReqChan: make(chan worker_v2.WorkRequest),
	}
	p.Resize(n)
	return p
}

func (p *Pool) Resize(n int){
	p.Mu.Lock()
	defer p.Mu.Unlock()
	n_workers:=len(p.Workers)
	if n_workers==n{
		return
	}
	for i:=n_workers;i<n ;i++  {
		p.Workers=append(p.Workers,worker_v2.NewWorker(i,p.ReqChan))
	}
	for i:=n;i<n_workers;i++{
		p.Workers[i].Stop()
	}
	p.Workers=p.Workers[0:n]
}

func (p *Pool)EnQueue(job worker_v2.Job){
	work:=<-p.ReqChan
	work.JobChan<-job
}

func (p *Pool)WorkerCount()int{
	return len(p.Workers)
}

