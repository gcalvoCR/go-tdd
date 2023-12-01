package routes

import (
	"encoding/json"
	"net/http"
)

type health struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	h := health{
		Status:   "OK",
		Messages: []string{},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h)
}
