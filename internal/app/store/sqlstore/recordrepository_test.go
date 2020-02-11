package sqlstore_test

import (
	"strconv"
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
	"github.com/JohnNON/gamewithnums/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestRecordRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users", "records")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

	r := model.TestRecord(t)
	r.UserID = u.ID

	assert.NoError(t, s.Record().Create(r))
	assert.NotNil(t, r)
}

func TestRecordRepository_FindByUserID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users", "records")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	r := model.TestRecord(t)
	_, err := s.Record().FindByUserID("1", "1")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	r.UserID = u.ID
	s.Record().Create(r)

	rc, err := s.Record().FindByUserID(strconv.Itoa(u.ID), strconv.Itoa(r.Difficulty))
	assert.NoError(t, err)
	assert.NotNil(t, rc)

}

func TestRecordRepository_GetAllRecords(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseDriver, databaseURL)
	defer teardown("users", "records")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	r := model.TestRecord(t)
	_, err := s.Record().GetAllRecords("1")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	r.UserID = u.ID
	s.Record().Create(r)

	rc, err := s.Record().GetAllRecords(strconv.Itoa(r.Difficulty))
	assert.NoError(t, err)
	assert.NotNil(t, rc)

}
