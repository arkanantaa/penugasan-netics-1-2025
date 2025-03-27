package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResp struct {
	Nama      string    `json:"nama"`
	NRP       string    `json:"nrp"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)

	currtime := time.Now()

	response := HealthResp{
		Nama:      "Arkananta Masarief",
		NRP:       "5025231115",
		Status:    "UP",
		Timestamp: currtime,
		Uptime:    uptime.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", handlerHealth)

	http.ListenAndServe(":8080", nil)
}
