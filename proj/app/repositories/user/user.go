package user

import (
	"fmt"
	"otus-arch-hw/app/entities"
	"otus-arch-hw/app/system/database/psql"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	InsertUser(user entities.User) (string, error)
	UpdateUser(user entities.User) error
	DeleteUser(emails []string) error
	GetUsers() ([]*entities.User, error)
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

func (r *repositoryDB) UpdateUser(user entities.User) error {
	query, args, err := prepareUpdateUserQuery(user)
	if err != nil {
		return fmt.Errorf("failed to prepare update user query: %w", err)
	}

	if _, err := r.dbClient.Exec(query, args...); err != nil {
		return fmt.Errorf("failed to update user query: %w", err)
	}

	return nil
}

func (r *repositoryDB) DeleteUser(emails []string) error {
	query, args, err := prepareDeleteUserQuery(emails)
	if err != nil {
		return fmt.Errorf("failed to prepare delete user query: %w", err)
	}

	if _, err := r.dbClient.Exec(query, args...); err != nil {
		return fmt.Errorf("failed to delete user query: %w", err)
	}

	return nil
}

func (r *repositoryDB) GetUsers() ([]*entities.User, error) {
	query, args, err := prepareGetUsersQuery()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get users query: %w", err)
	}

	rows, err := r.dbClient.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get users query: %w", err)
	}
	defer rows.Close()

	var result []*entities.User
	for rows.Next() {
		var target entities.User
		if err := rows.StructScan(&target); err != nil {
			continue
		}

		result = append(result, &target)
	}

	return result, nil
}
