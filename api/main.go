package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func dbReachable() bool {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"app":         "idp-test-app-multiple-api",
			"dbReachable": dbReachable(),
		})
	})
	addr := ":8080"
	fmt.Println("api listening on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("server error:", err)
		os.Exit(1)
	}
}
