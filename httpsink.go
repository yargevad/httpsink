package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port *int = flag.Int("port", 80, "port to listen on")
var https *bool = flag.Bool("https", false, "listen for https connections")
var cert *string = flag.String("cert", "", "path to https cert")
var key *string = flag.String("key", "", "path to https key")

func noOp(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	switch req.Method {
	case "PUT", "POST":
		_, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("ERROR: reading request data: %s\n", err)
		}
	}
	return
}

func main() {
	flag.Parse()

	// sanity check options
	if port == nil {
		log.Fatal("no port found")
	}
	if https != nil && *https {
		if cert != nil && *cert == "" {
			log.Fatal("--cert is required when using --https")
		} else if key != nil && *key == "" {
			log.Fatal("--key is required when using --https")
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", noOp)
	portSpec := fmt.Sprintf(":%d", *port)
	if https != nil && *https {
		log.Fatal(http.ListenAndServeTLS(portSpec, *cert, *key, mux))
	} else {
		log.Fatal(http.ListenAndServe(portSpec, mux))
	}
}
