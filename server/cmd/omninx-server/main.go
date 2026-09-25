package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type healthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Stage   string `json:"stage"`
}

func main() {
	addr := os.Getenv("OMNINX_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(healthResponse{
			Service: "OmniNX Server",
			Status:  "ok",
			Stage:   "foundation",
		})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		_, _ = w.Write([]byte("OmniNX server foundation: protocol services not implemented yet\n"))
	})

	log.Printf("OmniNX server foundation listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
