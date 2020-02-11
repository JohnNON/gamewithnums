package teststore_test

import (
	"strconv"
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
	"github.com/JohnNON/gamewithnums/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestRoundRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))

	rn := model.TestRound(t)
	rn.UserID = u.ID
	s.Round().Create(rn)
	assert.NoError(t, s.Round().Create(rn))

	assert.NotNil(t, rn)
}

func TestRoundRepository_FindByUserID(t *testing.T) {
	s := teststore.New()

	_, err := s.Round().FindByUserID("1")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)
	rn := model.TestRound(t)
	rn.UserID = u.ID
	s.Round().Create(rn)

	rounds, err := s.Round().FindByUserID(strconv.Itoa(u.ID))
	assert.NoError(t, err)
	assert.NotNil(t, rounds)

}

func TestRecordRepository_DeleteByUserID(t *testing.T) {
	s := teststore.New()

	_, err := s.Round().FindByUserID("1")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)
	rn := model.TestRound(t)
	rn.UserID = u.ID
	s.Round().Create(rn)
	assert.NoError(t, s.Round().DeleteByUserID(strconv.Itoa(u.ID)))
	rounds, err := s.Round().FindByUserID(strconv.Itoa(u.ID))
	assert.Error(t, err)
	assert.Nil(t, rounds)

}
