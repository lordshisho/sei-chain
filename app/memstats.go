package app

import (
	"encoding/json"
	"net/http"
	"runtime"
)

func getMemStats(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memStats)
}

func startMemStats() {
	http.HandleFunc("/memstats", getMemStats)

	// Start the HTTP server
	go func() {
		http.ListenAndServe(":6161", nil)
	}()
}
