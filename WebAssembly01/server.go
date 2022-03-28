// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	err := http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))
	log.Fatalln(err)
}
