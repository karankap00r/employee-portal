package server

import (
	"net/http"

	"github.com/karankap00r/employee-portal/config"
)

func Start(cfg config.Config) {
	http.HandleFunc("/cercli/", handleUsers)

	err := http.ListenAndServe(cfg.ServerAddress, nil)
	if err != nil {
		return
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	// Handle users endpoint
}
