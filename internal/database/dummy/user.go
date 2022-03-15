package dummy

import (
	"sync"
	"time"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/alfreddobradi/heymark/internal/helper"
	"github.com/google/uuid"
)

type UserRepository struct {
	mx      *sync.Mutex
	Records map[uuid.UUID]model.User
}

func NewUserRepository() UserRepository {
	return UserRepository{
		mx:      &sync.Mutex{},
		Records: map[uuid.UUID]model.User{},
	}
}

func (db *DummyDB) Authenticate(username, password string) (model.Token, error) {
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	for i := range db.Users.Records {
		user := db.Users.Records[i]
		if user.Username == username && user.Password == helper.Sha256(password) {
			token := model.NewToken()
			user.Tokens = append(user.Tokens, token) // TODO Clean periodically
			db.Users.Records[i] = user
			return token, nil
		}
	}

	return model.Token{}, model.ErrUserNotFound
}

func (db *DummyDB) authorize(auth model.AuthData) (model.User, error) {
	for _, user := range db.Users.Records {
		if user.Username == auth.Username {
			for _, t := range user.Tokens {
				if t.ID == auth.TokenID && t.ValidUntil.After(time.Now()) {
					return user, nil
				}
			}
		}
	}

	return model.User{ID: uuid.Nil}, model.ErrUnauthorized
}

func (db *DummyDB) Authorize(auth model.AuthData) (model.User, error) {
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	return db.authorize(auth)
}

func (db *DummyDB) GetUser(id uuid.UUID) (model.User, error) {
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	if user, ok := db.Users.Records[id]; ok {
		return user, nil
	}

	return model.User{}, model.ErrUserNotFound
}

func (db *DummyDB) CreateUser(username, password, bio string) (model.User, error) {
	db.Users.mx.Lock()
	defer db.Users.mx.Unlock()

	for _, user := range db.Users.Records {
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
		Bio:      bio,
	}

	db.Users.Records[id] = user

	return user, nil
}
