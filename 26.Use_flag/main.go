package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I try to use %s\n", r.URL.Path[1:])
}

func pinger(port string) error {
	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Errorf("server's status is not 200")
	}
	return nil

}

func main() {
	var port string
	var ping bool
	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&port, "p", "8080", "server port")
	flag.BoolVar(&ping, "ping", false, "ping server is live")
	flag.Parse()

	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	if ping {
		if err := pinger(port); err != nil {
			log.Printf("Server had some error")
			return
		}
		log.Printf("Server is live")
		return
	}

	http.HandleFunc("/", handler)
	log.Println("http server port is " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
