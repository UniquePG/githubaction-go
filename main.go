package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", RootHandler)

	http.HandleFunc("/hello", HelloHandler)

	fmt.Println("server started...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from server")

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "this is the hello route running from docker")

}