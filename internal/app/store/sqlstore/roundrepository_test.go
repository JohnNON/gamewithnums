package sqlstore_test

import (
	"strconv"
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
	"github.com/JohnNON/gamewithnums/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestRoundRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users", "rounds")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

	r := model.TestRound(t)
	r.UserID = u.ID

	assert.NoError(t, s.Round().Create(r))
	assert.NotNil(t, r)
}

func TestRoundRepository_FindByUserID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users", "rounds")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	r := model.TestRound(t)
	_, err := s.Round().FindByUserID("1")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	r.UserID = u.ID
	s.Round().Create(r)

	rc, err := s.Round().FindByUserID(strconv.Itoa(u.ID))
	assert.NoError(t, err)
	assert.NotNil(t, rc)

}

func TestRoundRepository_DeleteByUserID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users", "rounds")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	r := model.TestRound(t)
	_, err := s.Round().FindByUserID(strconv.Itoa(u.ID))
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	r.UserID = u.ID
	s.Round().Create(r)
	s.Round().Create(r)

	err = s.Round().DeleteByUserID(strconv.Itoa(u.ID))
	assert.NoError(t, err)

	_, err = s.Round().FindByUserID(strconv.Itoa(u.ID))
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

}
