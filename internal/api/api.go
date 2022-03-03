package api

import (
	"io"
	"log"
)

func Execute(body io.Reader) ([]byte, error) {
	raw, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	log.Println(string(raw))
	return []byte("{}"), nil
}
