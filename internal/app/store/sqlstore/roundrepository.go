package sqlstore

import (
	"database/sql"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// RoundRepository - способ хранения user
type RoundRepository struct {
	store *Store
}

// Create - создаст место хранения user
func (r *RoundRepository) Create(rn *model.Round) error {
	if err := rn.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO rounds (userid, difficulty, gamenumber, gametime, inpt, outpt) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		rn.UserID,
		rn.Difficulty,
		rn.GameNumber,
		rn.GameTime,
		rn.Inpt,
		rn.Outpt,
	).Scan(&rn.ID)
}

// FindByUserID - ищет round по значению поля userId
func (r *RoundRepository) FindByUserID(userID string) (*[]model.Round, error) {
	rn := &[]model.Round{}
	if err := r.store.db.Select(
		rn,
		`SELECT difficulty, gamenumber, gametime, inpt, outpt FROM rounds
		WHERE userid = $1 ORDER BY id`,
		userID,
	); err != nil || len(*rn) == 0 {
		if err == sql.ErrNoRows || len(*rn) == 0 {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return rn, nil
}

// DeleteByUserID - очищает записи по userID
func (r *RoundRepository) DeleteByUserID(userID string) error {
	if _, err := r.store.db.Exec(
		`DELETE FROM rounds WHERE userid = $1`,
		userID,
	); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}

	return nil
}

// RoundCheck - проверяет есть ли записи с userID
func (r *RoundRepository) RoundCheck(userID int) bool {
	rnd := &model.Round{}
	if err := r.store.db.QueryRow(
		"SELECT id FROM rounds WHERE userid = $1 LIMIT 1",
		userID,
	).Scan(
		&rnd.ID,
	); err != nil {

		return false
	}

	return true
}
