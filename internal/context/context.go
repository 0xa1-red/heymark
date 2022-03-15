package context

import (
	"context"
	"fmt"

	"github.com/alfreddobradi/heymark/internal/database/model"
)

type Auth string

var (
	ErrNoAuthorization      error = fmt.Errorf("No authorization header found")
	ErrInvalidAuthorization error = fmt.Errorf("Invalid authorization header")
)

func GetAuthData(ctx context.Context) (model.AuthData, error) {
	authHeader := ctx.Value(Auth("Authorization"))
	if authHeader == nil {
		return model.AuthData{}, ErrNoAuthorization
	}

	header, ok := authHeader.(string)
	if !ok {
		return model.AuthData{}, ErrInvalidAuthorization
	}

	return model.GetAuthDataFromHeader(header)
}
