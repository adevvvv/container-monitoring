package repository

import (
	"container-monitoring/backend/model"
	"database/sql"
	"fmt"
)

type PingRepository interface {
	GetAll() ([]model.PingStatus, error)
	GetByID(id int) (*model.PingStatus, error)
	Create(status *model.PingStatus) error
	Update(status *model.PingStatus) error
	Delete(id int) error
}

type PostgresPingRepository struct {
	db *sql.DB
}

func NewPostgresPingRepository(db *sql.DB) *PostgresPingRepository {
	return &PostgresPingRepository{db: db}
}

func (r *PostgresPingRepository) GetAll() ([]model.PingStatus, error) {
	rows, err := r.db.Query("SELECT id, ip, ping_time, last_success FROM ping_status ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []model.PingStatus
	for rows.Next() {
		var status model.PingStatus
		if err := rows.Scan(&status.ID, &status.IP, &status.PingTime, &status.LastSuccess); err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
}

func (r *PostgresPingRepository) GetByID(id int) (*model.PingStatus, error) {
	var status model.PingStatus
	err := r.db.QueryRow("SELECT id, ip, ping_time, last_success FROM ping_status WHERE id = $1", id).
		Scan(&status.ID, &status.IP, &status.PingTime, &status.LastSuccess)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("status with id %d not found", id)
		}
		return nil, err
	}
	return &status, nil
}

func (r *PostgresPingRepository) Create(status *model.PingStatus) error {
	err := r.db.QueryRow(
		"INSERT INTO ping_status (ip, ping_time, last_success) VALUES ($1, $2, $3) RETURNING id",
		status.IP, status.PingTime, status.LastSuccess,
	).Scan(&status.ID)
	return err
}

func (r *PostgresPingRepository) Update(status *model.PingStatus) error {
	result, err := r.db.Exec(
		"UPDATE ping_status SET ip = $1, ping_time = $2, last_success = $3 WHERE id = $4",
		status.IP, status.PingTime, status.LastSuccess, status.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated for status with id %d", status.ID)
	}
	return nil
}

func (r *PostgresPingRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM ping_status WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows deleted for status with id %d", id)
	}
	return nil
}
