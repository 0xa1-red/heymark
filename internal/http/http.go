package http

import (
	"log"
	"net/http"

	"github.com/alfreddobradi/heymark/internal/api"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
}

func New(addr string) *Server {
	handler := mux.NewRouter()

	handler.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s /api", r.Method)
		response, err := api.Execute(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(response) // nolint
	})

	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	spew.Dump(s)

	return &Server{s}
}
