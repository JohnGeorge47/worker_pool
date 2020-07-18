package worker_pool

import (
	"fmt"
	"sync"
)

type Job struct {
	Id  int
	Val string
}

type Pool struct {
	Mu   sync.Mutex
	Size int
	Kill chan struct{}
	Jobs chan Job
	wg    sync.WaitGroup
}

func NewPool(size int) *Pool {
	pool := &Pool{
		Jobs: make(chan Job, 100),
		Kill: make(chan struct{}),
	}
	pool.Resize(size)
	return pool
}
func (p *Pool) Resize(n int) {
	fmt.Println(p.Size,n)
	p.Mu.Lock()
	defer p.Mu.Unlock()
	if p.Size<n{
		for p.Size < n {
			p.Size++
			p.wg.Add(1)
			go p.worker(p.Size)
		}
	}else {
		for p.Size > n {
			p.Size--
			p.Kill <- struct{}{}
		}
	}
}


func (p *Pool) worker(id int) {
	defer p.wg.Done()
	fmt.Println("Starting worker id:",id)
	for {
		select {
		case j, ok := <-p.Jobs:
			if !ok {
				return
			}
			fmt.Println("worker no:",id)
			j.Execute()
		case <-p.Kill:
			fmt.Println("Killing worker:",id)
			return
		}
	}
}

func (p *Pool) Wait() {
	p.wg.Wait()
}


func (p Pool) AddTask(job Job) {
	fmt.Println("here1")
	p.Jobs <- job
}

func (j Job) Execute() {
	fmt.Printf("Job id is %d Job value is %s", j.Id,j.Val)
}
