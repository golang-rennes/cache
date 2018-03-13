package main

import (
	"cache"
	"fmt"
	"log"
	"net/http"
)

const port = 8888

func main() {
	http.Handle("/", reverse{cache: cache.New()})

	portStr := fmt.Sprintf(":%d", port)
	log.Print("Server listening on http://localhost" + portStr)
	log.Fatal(http.ListenAndServe(portStr, nil))
}
