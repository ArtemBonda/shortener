package app

import (
	"fmt"
	"net/http"
)

func RootAcceptURL(wr http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(wr, "Method GET")
	case http.MethodPost:
		fmt.Fprintln(wr, "Method POST")
	default:
		http.Error(wr, "not found", http.StatusNotFound)
	}
}
