package main

import (
	"fmt"
	"go.elastic.co/apm/module/apmgorilla"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)
	r.Use(apmgorilla.Middleware())
	log.Fatal(http.ListenAndServe(":8000", r))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
}
