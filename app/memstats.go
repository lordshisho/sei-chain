package app

import (
	"encoding/json"
	"net/http"
	"runtime"
	"runtime/debug"
)

func getMemStats(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memStats)
}

func release(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	debug.FreeOSMemory()

	// get memstats and release
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
}

func startMemStats() {
	http.HandleFunc("/memstats", getMemStats)
	http.HandleFunc("/release", release)

	// Start the HTTP server
	go func() {
		http.ListenAndServe(":6161", nil)
	}()
}
