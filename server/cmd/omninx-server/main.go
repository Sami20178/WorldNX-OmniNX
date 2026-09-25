package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go/v2"
)

type healthResponse struct {
	Service string
	Status  string
	Stage   string
	PRUDP   string
}

func envInt(name string, fallback int) int {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed < 1 || parsed > 65535 {
		return fallback
	}
	return parsed
}

func main() {
	httpAddr := os.Getenv("OMNINX_HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = "127.0.0.1:8080"
	}

	prudpPort := envInt("OMNINX_PRUDP_PORT", 60000)

	prudpServer := nex.NewPRUDPServer()
	prudpServer.SetFragmentSize(962)
	prudpServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(1, 1, 0))
	prudpServer.SessionKeyLength = 16

	// Initial OmniNX service endpoint.
	// Authentication and game-specific handlers are added by service modules.
	endpoint := nex.NewPRUDPEndPoint(1)
	prudpServer.BindPRUDPEndPoint(endpoint)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(healthResponse{
			Service: "OmniNX Server",
			Status:  "ok",
			Stage:   "prudp-foundation",
			PRUDP:   ":" + strconv.Itoa(prudpPort),
		})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OmniNX Server is running\n"))
	})

	go func() {
		log.Printf("OmniNX HTTP health server listening on %s", httpAddr)
		if err := http.ListenAndServe(httpAddr, mux); err != nil {
			log.Printf("HTTP server stopped: %v", err)
		}
	}()

	log.Printf("OmniNX PRUDP server listening on UDP :%d", prudpPort)
	if err := prudpServer.Listen(prudpPort); err != nil {
		log.Fatalf("PRUDP server failed: %v", err)
	}
}
