package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello "+ r.URL.Path[1:] + "!\n")
}

func main() {
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("server")
}
