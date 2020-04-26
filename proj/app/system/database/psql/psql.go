package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	HealthCheck() bool
	GetConnection() *sqlx.DB
}

type repository struct {
	Client *sqlx.DB
}

func New(dsn string) (Repository, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	repo := &repository{Client: db}
	if err = repo.Client.Ping(); err != nil {
		return repo, err
	}

	return repo, nil
}

func (db *repository) GetConnection() *sqlx.DB {
	return db.Client
}

func (db *repository) HealthCheck() bool {
	if err := db.Client.Ping(); err != nil {
		return false
	}

	return true
}
