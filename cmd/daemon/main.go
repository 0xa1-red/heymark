package main

import (
	"log"

	"github.com/alfreddobradi/heymark/internal/http_service"
)

func main() {
	s := http_service.New("localhost:9456")

	if err := s.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
