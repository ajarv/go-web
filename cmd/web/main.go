package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// health handler
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"UP"}`)
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func main() {
	var host string
	var port string
	flag.StringVar(&host, "host", getEnv("LISTEN_ADDRESS", "0.0.0.0"), "listen host")
	flag.StringVar(&port, "port", getEnv("SERVICE_PORT", "8080"), "listen port")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    host + ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Fprintf(os.Stdout, "Server listening HTTP %s:%s\n", host, port)
	log.Fatal(srv.ListenAndServe())
}
