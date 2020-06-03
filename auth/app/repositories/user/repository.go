package user

import (
	"fmt"
	"otus-auth/app/entities"
	"otus-auth/app/system/database/psql"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	InsertUser(user entities.User) (string, error)
	UpdateUser(userEmail string, newUserInfo entities.User) error
	GetUser(email string) (*entities.User, error)
}

type repositoryDB struct {
	dbClient *sqlx.DB
}

func NewRepository(db psql.Repository) *repositoryDB {
	return &repositoryDB{dbClient: db.GetConnection()}
}

func (r *repositoryDB) InsertUser(user entities.User) (userID string, err error) {
	query, args, err := prepareInsertUserQuery(user)
	if err != nil {
		return userID, fmt.Errorf("failed to prepare insert user query: %w", err)
	}

	res := r.dbClient.QueryRowx(query, args...)
	if err = res.Scan(&userID); err != nil {
		return userID, fmt.Errorf("failed to scan result before insert user query: %w", err)
	}

	return userID, nil
}

func (r *repositoryDB) UpdateUser(userEmail string, newUserInfo entities.User) error {
	query, args, err := prepareUpdateUserQuery(userEmail, newUserInfo)
	if err != nil {
		return fmt.Errorf("failed to prepare update user query: %w", err)
	}

	if _, err := r.dbClient.Exec(query, args...); err != nil {
		return fmt.Errorf("failed to update user query: %w", err)
	}

	return nil
}

func (r *repositoryDB) GetUser(email string) (*entities.User, error) {
	query, args, err := prepareGetUserQuery(email)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get user query: %w", err)
	}

	user := &entities.User{}
	row := r.dbClient.QueryRowx(query, args...)
	if err := row.StructScan(user); err != nil {
		return nil, err
	}

	return user, nil
}
