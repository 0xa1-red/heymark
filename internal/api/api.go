package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	int_context "github.com/alfreddobradi/heymark/internal/context"
	"github.com/alfreddobradi/heymark/internal/schema"
	"github.com/graphql-go/graphql"
)

func Execute(r *http.Request) ([]byte, error) {
	decoder := json.NewDecoder(r.Body)

	var params map[string]interface{}
	err := decoder.Decode(&params)
	if err != nil {
		return nil, err
	}

	if query, ok := params["query"]; ok {
		ctx := r.Context()
		if auth := r.Header.Get("Authorization"); auth != "" {
			ctx = context.WithValue(ctx, int_context.Auth("Authorization"), auth)
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema.RootSchema,
			RequestString: query.(string),
			Context:       ctx,
		})

		b := bytes.NewBuffer([]byte{})
		encoder := json.NewEncoder(b)
		err := encoder.Encode(result)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal response: %w", err)
		}

		return b.Bytes(), nil
	}

	// log.Printf("%+v", params)
	return nil, fmt.Errorf("Not found")
}
