package main

import (
	"net/http"
	"encoding/json"
)

type HealthCheck struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		data := &HealthCheck{Status:"OK"}
		response, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
