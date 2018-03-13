package main

import (
	"cache"
	"net/http"
	"strings"
	"time"
)

type reverse struct {
	cache cache.T
}

// ServeHTTP handles HTTP requests, and implements the http.Handler interface
func (r reverse) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.Trim(req.RequestURI, " /")
	resp, ok := r.cache.Get(path)
	if !ok {
		resp = computeResponse(path)
		r.cache.Add(path, resp)
	}

	w.Write(resp)
}

// computeResponse simulates a long computation…
func computeResponse(s string) []byte {
	time.Sleep(time.Second)

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
