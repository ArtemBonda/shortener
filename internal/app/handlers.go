package app

import (
	"fmt"
	"github.com/ArtemBonda/shortener/internal/entity"
	"github.com/ArtemBonda/shortener/internal/hashing"
	"io"
	"net/http"
	"net/url"
)

var DB = entity.Short{}

func RootAcceptURL(wr http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Path[1:]

		originURL, err := DB.SearchLink(id)
		if err != nil {
			fmt.Println("not found")
			http.Error(wr, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Redirect(wr, r, originURL, http.StatusTemporaryRedirect)

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(wr, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if _, err = url.ParseRequestURI(string(body)); err != nil {
			http.Error(wr, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		_, err = http.Get(string(body))
		if err != nil {
			http.Error(wr, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		key := hashing.HashURLAddr(string(body))
		fmt.Println(string(body))
		err = DB.AddLink(key, string(body))
		if err != nil {
			http.Error(wr, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
			return
		}
		wr.WriteHeader(http.StatusCreated)
	default:
		http.Error(wr, "not found", http.StatusNotFound)
	}
}
