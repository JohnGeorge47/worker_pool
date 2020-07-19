package main

import (
	"flag"
	"fmt"
	"github.com/JohnGeorge47/sendx/handlers"
	"github.com/JohnGeorge47/sendx/pkg/dispatcher"
	"github.com/gorilla/mux"
	"net/http"
)

var port = flag.String("port", "3003", "You know the port or something")

func main() {
    d:=dispatcher.StartDispatcher(10)
	r := mux.NewRouter()
	r.Handle("/produce", handlers.Collector(d))
	fmt.Printf("Listening on port %s",*port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
	if err != nil {
		fmt.Println(err)
	}
}
