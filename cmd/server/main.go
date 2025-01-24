package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go func() {
		// Serve asserts on port 9090
		if err := http.ListenAndServe(":9090", http.FileServer(http.Dir("assets"))); err != nil {
			panic(fmt.Errorf("Error starting server on port 9090: %w", err))
		}
	}()

	// Simulate a slow response on port 9091
	if err := http.ListenAndServe(":9091", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintln(w, "ok")
	})); err != nil {
		panic(fmt.Errorf("Error starting server on port 9091: %w", err))

	}
}
