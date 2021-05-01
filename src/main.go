package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
}

var (
	okStatus     = &HealthResponse{Status: "ok"}
	okStatusJson string
)

func init() {
	temp, _ := json.Marshal(okStatus)
	okStatusJson = string(temp)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "Hello from home")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, okStatusJson)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/health", HealthHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server")
	log.Fatal(srv.ListenAndServe())
}
