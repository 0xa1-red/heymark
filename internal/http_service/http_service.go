package http_service

import (
	"log"
	"net/http"

	"github.com/alfreddobradi/heymark/internal/api"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
}

func New(addr string) *Server {
	handler := mux.NewRouter()

	handler.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s /api", r.Method)
		response, err := api.Execute(r)
		if err != nil {
			log.Printf("Error: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response) // nolint
	})

	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server{s}
}
