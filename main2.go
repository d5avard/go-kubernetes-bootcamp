package main

// import (
// 	"log"
// 	"math"
// 	"net/http"
// 	"os"
// 	"time"
// )

// var requests int = 0
// var startTime time.Time = time.Now().Local()

// func main() {

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", handleIndex)
// 	// mux.HandleFunc("/cm", handleConfigMap)
// 	// mux.HandleFunc("/secrets", handleSecrets)

// 	server := &http.Server{Addr: ":8080", Handler: mux}

// 	log.Printf("Kubernetes Bootcamp App in GO Started At %s\n", startTime.Format(time.Stamp))
// 	if hostname, err := os.Hostname(); err != nil {
// 		log.Printf("Running on port: %s", hostname)
// 	}

// 	server.ListenAndServe()
// }

// func handleIndex(w http.ResponseWriter, req *http.Request) {
// 	header := w.Header()
// 	header.Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Hello Kubernetes bootcamp in GO! | Running on: "))
// 	w.Write([]byte(req.Host))
// 	w.Write([]byte(" | v=1\n"))

// 	requests++
// 	newTime := time.Now().Local()
// 	timeDiff := newTime.Sub(startTime) / time.Duration(math.Pow10(9))
// 	log.Printf("Running On: %s | Total Requests: %d | App Uptime: %d seconds | Log Time %s\n", req.Host, requests, timeDiff, newTime.Format(time.Stamp))
// }

// func handleConfigMap(w http.ResponseWriter, req *http.Request) {
// 	header := w.Header()
// 	header.Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)

// 	// ENV_APP_NAME
// 	var appName string = os.Getenv("ENV_APP_NAME")

// 	w.Write([]byte("Handle ConfigMap\n"))
// 	w.Write([]byte(appName))
// 	w.Write([]byte("/n"))
// }

// func handleSecrets(w http.ResponseWriter, req *http.Request) {
// 	header := w.Header()
// 	header.Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)

// 	w.Write([]byte("Handle Secrets\n"))
// }
