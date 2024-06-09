package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in Golang!")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello, this is a greeter function!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is a response to a http request</h1>"))
}
