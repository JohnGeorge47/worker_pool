package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"github.com/JohnGeorge47/sendx/pkg/worker_pool_v2"
	"github.com/JohnGeorge47/sendx/pkg/worker_v2"
)

type SendXRequest struct {
	Val    *string `json:"val"`
	Resize *int    `json:"resize"`
}

func Collector(p *worker_pool_v2.Pool) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req SendXRequest
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}
			json.Unmarshal(body, &req)
			rand.Seed(time.Now().UnixNano())
			if req.Resize != nil {
				fmt.Println(*req.Resize)
				fmt.Printf("Current number of workers %d",p.WorkerCount())
				p.Resize(*req.Resize)
				fmt.Printf("New worker count %d",p.WorkerCount())
			}
			if req.Val != nil {
				job := worker_v2.Job{
					Id: rand.Intn(100),
					Value: *req.Val,
				}
				go p.EnQueue(job)
				w.Write([]byte("A dummy response"))
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Only post method allowed"))
		}
	}
	return http.HandlerFunc(fn)
}
