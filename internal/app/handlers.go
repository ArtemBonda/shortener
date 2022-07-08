package app

import (
	"fmt"
	"github.com/ArtemBonda/shortener/internal/entity"
	"github.com/ArtemBonda/shortener/internal/hashing"
	"io"
	"net/http"
)

var DB = entity.Short{}

func RootAcceptURL(wr http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Path[1:]
		if id == "" {
			fmt.Println("URL not found")
			http.Error(wr, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		originURL, err := DB.SearchLink(id)
		if err != nil {
			fmt.Println("URL not found")
			http.Error(wr, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		wr.Header().Set("Location", originURL)
		wr.WriteHeader(http.StatusTemporaryRedirect)
		wr.Write([]byte(id))
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		key := hashing.HashURLAddr(string(body))
		fmt.Println(string(body))
		err = DB.AddLink(key, string(body))
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		wr.WriteHeader(http.StatusCreated)
		fmt.Fprintf(wr, "key: %s, value: %s", key, string(body))
	default:
		http.Error(wr, "not found", http.StatusNotFound)
	}
}
