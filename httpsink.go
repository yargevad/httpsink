package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port *int = flag.Int("port", 80, "port to listen on")

func noOp(w http.ResponseWriter, req *http.Request) {
	return
}

func main() {
	flag.Parse()
	if port == nil {
		log.Fatal("no port found")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", noOp)
	portSpec := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(portSpec, mux))
}
