package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the world of modules!!!")

	greet()

	r := mux.NewRouter()
	r.HandleFunc("/", serveWeb).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func greet() {
	fmt.Println("hey there mod users")

}

func serveWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))
}
