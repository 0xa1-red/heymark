package main

import (
	"log"

	"github.com/alfreddobradi/heymark/internal/http"
)

func main() {
	s := http.New("localhost:9000")

	if err := s.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
