package sqlstore

import (
	"database/sql"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// RecordRepository - способ хранения user
type RecordRepository struct {
	store *Store
}

// Create - создаст место хранения user
func (r *RecordRepository) Create(rc *model.Record) error {
	if err := rc.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO records (userid, difficulty, roundcount, gametime) VALUES ($1, $2, $3, $4) RETURNING id",
		rc.UserID,
		rc.Difficulty,
		rc.RoundCount,
		rc.GameTime,
	).Scan(&rc.ID)
}

// FindByUserID - ищет record по значению поля userId и difficulty
func (r *RecordRepository) FindByUserID(userID string, diff string) (*[]model.Record, error) {
	rc := &[]model.Record{}
	if err := r.store.db.Select(
		rc,
		"SELECT * FROM records WHERE userid = $1 AND difficulty = $2 ORDER BY gametime, roundcount",
		userID,
		diff,
	); err != nil || len(*rc) == 0 {
		if err == sql.ErrNoRows || len(*rc) == 0 {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return rc, nil
}

// GetAllRecords - ищет user по значению поля email
func (r *RecordRepository) GetAllRecords(diff string) (*[]model.Record, error) {
	rc := &[]model.Record{}
	if err := r.store.db.Select(
		rc,
		"SELECT * FROM records WHERE difficulty = $1 ORDER BY gametime, roundcount LIMIT 10",
		diff,
	); err != nil || len(*rc) == 0 {
		if err == sql.ErrNoRows || len(*rc) == 0 {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return rc, nil
}
