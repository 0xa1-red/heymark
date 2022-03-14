package dummy

import (
	"log"
	"time"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/alfreddobradi/heymark/internal/helper"
	"github.com/google/uuid"
)

func (db *DummyDB) Authenticate(username, password string) (model.Token, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	for i := range db.Users {
		user := db.Users[i]
		if user.Username == username && user.Password == helper.Sha256(password) {
			token := model.NewToken()
			user.Tokens = append(user.Tokens, token) // TODO Clean periodically
			db.Users[i] = user
			log.Printf("Tokens for %s: %+v", user.Username, user.Tokens)
			return token, nil
		}
	}

	return model.Token{}, model.ErrUserNotFound
}

func (db *DummyDB) Authorize(auth model.AuthData) (model.User, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	for _, user := range db.Users {
		if user.Username == auth.Username {
			for _, t := range user.Tokens {
				if t.ID == auth.TokenID && t.ValidUntil.After(time.Now()) {
					return user, nil
				}
			}
		}
	}

	return model.User{}, model.ErrUnauthorized
}

func (db *DummyDB) GetUser(id uuid.UUID) (model.User, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	if user, ok := db.Users[id]; ok {
		return user, nil
	}

	return model.User{}, model.ErrUserNotFound
}

func (db *DummyDB) CreateUser(username, password string) (model.User, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	for _, user := range db.Users {
		if user.Username == username {
			return model.User{}, model.ErrUsernameExists
		}
	}

	id := uuid.New()
	passHash := helper.Sha256(password)

	user := model.User{
		ID:       id,
		Username: username,
		Password: passHash,
	}

	db.Users[id] = user

	return user, nil
}
