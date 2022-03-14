package dummy

import (
	"sync"
	"time"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/alfreddobradi/heymark/internal/helper"
	"github.com/google/uuid"
)

type userRepository struct {
	mx      *sync.Mutex
	records map[uuid.UUID]model.User
}

func (db *DummyDB) Authenticate(username, password string) (model.Token, error) {
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	for i := range db.Users.records {
		user := db.Users.records[i]
		if user.Username == username && user.Password == helper.Sha256(password) {
			token := model.NewToken()
			user.Tokens = append(user.Tokens, token) // TODO Clean periodically
			db.Users.records[i] = user
			return token, nil
		}
	}

	return model.Token{}, model.ErrUserNotFound
}

func (db *DummyDB) Authorize(auth model.AuthData) (model.User, error) {
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	for _, user := range db.Users.records {
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
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	if user, ok := db.Users.records[id]; ok {
		return user, nil
	}

	return model.User{}, model.ErrUserNotFound
}

func (db *DummyDB) CreateUser(username, password string) (model.User, error) {
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	for _, user := range db.Users.records {
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

	db.Users.records[id] = user

	return user, nil
}
