package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT string

func handleRequests(p string) {
	declareHandlers()
	fmt.Printf("Listening on 0.0.0.0:%s\n", p)
	log.Fatal(http.ListenAndServe(":"+p, nil))
}

func main() {
	PORT, e := os.LookupEnv("PORT")
	if !e {
		PORT = "8088"
	}
	handleRequests(PORT)
}
