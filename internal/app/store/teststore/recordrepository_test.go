package teststore_test

import (
	"strconv"
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
	"github.com/JohnNON/gamewithnums/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestRecordRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestRecordRepository_FindByUserID(t *testing.T) {
	s := teststore.New()

	_, err := s.Record().FindByUserID("1", "1")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)
	rc := model.TestRecord(t)
	rc.UserID = u.ID
	s.Record().Create(rc)

	records, err := s.Record().FindByUserID(strconv.Itoa(u.ID), strconv.Itoa(rc.Difficulty))
	assert.NoError(t, err)
	assert.NotNil(t, records)

}

func TestRecordRepository_GetAllRecords(t *testing.T) {
	s := teststore.New()

	_, err := s.Record().GetAllRecords("1")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)
	rc := model.TestRecord(t)
	rc.UserID = u.ID
	s.Record().Create(rc)

	records, err := s.Record().GetAllRecords(strconv.Itoa(rc.Difficulty))
	assert.NoError(t, err)
	assert.NotNil(t, records)

}
