package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

var requests int = 0
var startTime time.Time = time.Now().Local()

func main() {

	handleRequest := func(w http.ResponseWriter, req *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Kubernetes bootcamp in GO! | Running on: "))
		w.Write([]byte(req.Host))
		w.Write([]byte(" | v=1\n"))

		requests++
		newTime := time.Now().Local()
		timeDiff := newTime.Sub(startTime) / time.Duration(math.Pow10(9))
		log.Printf("Running On: %s | Total Requests: %d | App Uptime: %d seconds | Log Time %s\n", req.Host, requests, timeDiff, newTime.Format(time.Stamp))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)
	server := &http.Server{Addr: ":8080", Handler: mux}

	log.Printf("Kubernetes Bootcamp App in GO Started At %s\n", startTime.Format(time.Stamp))
	if hostname, err := os.Hostname(); err != nil {
		log.Printf("Running on port: %s", hostname)
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
