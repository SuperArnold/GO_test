package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I try to use %s\n", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("http server port is 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
