package http_service

import (
	"context"
	"log"
	"net/http"

	"github.com/alfreddobradi/heymark/internal/api"
	int_context "github.com/alfreddobradi/heymark/internal/context"
	"github.com/alfreddobradi/heymark/internal/database"
	"github.com/google/uuid"
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

	handler.HandleFunc("/bookmark/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		ctx := r.Context()
		if auth := r.Header.Get("Authorization"); auth != "" {
			ctx = context.WithValue(ctx, int_context.Auth("Authorization"), auth)
		}

		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		db := database.GetDB()
		bookmark, err := db.GetBookmark(ctx, id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, bookmark.URL, http.StatusPermanentRedirect)
	}).Methods(http.MethodGet)

	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server{s}
}
