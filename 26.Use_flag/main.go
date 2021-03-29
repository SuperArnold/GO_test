package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I try to use %s\n", r.URL.Path[1:])
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&port, "p", "8080", "server port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("http server port is " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
