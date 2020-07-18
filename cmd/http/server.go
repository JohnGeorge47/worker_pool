package main

import (
	"flag"
	"fmt"
	"github.com/JohnGeorge47/sendx/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/JohnGeorge47/sendx/pkg/worker_pool"
)

var port = flag.String("port", "3003", "You know the port or something")

func main() {
    pool:=worker_pool.NewPool(10)
	r := mux.NewRouter()
	r.Handle("/produce", handlers.Collector(pool))
	fmt.Printf("Listening on port %s",*port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
	if err != nil {
		fmt.Println(err)
	}
}
