package sqlstore_test

import (
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
	"github.com/JohnNON/gamewithnums/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	s.User().Create(u)

	u, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.com"
	u := model.TestUser(t)
	u.Email = email
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindByNickname(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	nickname := "user"
	u := model.TestUser(t)
	u.Nickname = nickname
	_, err := s.User().FindByNickname(nickname)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)

	u, err = s.User().FindByNickname(nickname)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
