package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/JohnGeorge47/sendx/pkg/worker_pool"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type SendXRequest struct {
	Val *string `json:"val"`
	Resize *int `json:"resize"`
}

func Collector(pool *worker_pool.Pool) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method==http.MethodPost{
			var req SendXRequest
			body, err := ioutil.ReadAll(r.Body)
			if err!=nil{
				fmt.Println(err)
			}
			json.Unmarshal(body,&req)
			rand.Seed(time.Now().UnixNano())
			if req.Resize!=nil{
				pool.Resize(*req.Resize)
			}
			if req.Val!=nil{
				task:=worker_pool.Job{
					Id:rand.Intn(100)  ,
					Val:*req.Val,
				}
				pool.AddTask(task)
				pool.Wait()
				fmt.Println(req)
			}
		}else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Only post method allowed"))
		}
	}
	return http.HandlerFunc(fn)
}
