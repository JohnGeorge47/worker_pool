package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/JohnGeorge47/sendx/pkg/dispatcher"
	"github.com/JohnGeorge47/sendx/pkg/worker"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type SendXRequest struct {
	Val    *string `json:"val"`
	Resize *int    `json:"resize"`
}

func Collector(dispatcher *dispatcher.Dispatcher) http.Handler {
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
				fmt.Println(dispatcher.Size)
				dispatcher.Resize(*req.Resize)
			}
			if req.Val != nil {
				job := worker.Job{
					JobId: rand.Intn(100),
					Value: *req.Val,
				}
				fmt.Println(req)
				dispatcher.JobQueue <- job
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Only post method allowed"))
		}
	}
	return http.HandlerFunc(fn)
}
