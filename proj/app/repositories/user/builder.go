package user

import (
	"otus-arch-hw/app/entities"

	"github.com/Masterminds/squirrel"
)

const tableUsers = "users"

func prepareInsertUserQuery(user entities.User) (string, []interface{}, error) {
	psqlSq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	rawRequest := psqlSq.Insert(tableUsers).
		Columns("email", "password_hash", "first_name", "last_name").
		Values(user.Email, user.PasswordHash, user.FirstName, user.LastName).
		Suffix("returning email")

	return rawRequest.ToSql()
}

func prepareUpdateUserQuery(user entities.User) (string, []interface{}, error) {
	psqlSq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	rawRequest := psqlSq.Update(tableUsers).
		SetMap(user.ToMap()).
		Where(squirrel.Eq{
			"email": user.Email,
		})

	return rawRequest.ToSql()
}

func prepareDeleteUserQuery(emails []string) (string, []interface{}, error) {
	psqlSq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	rawRequest := psqlSq.Delete(tableUsers).
		Where(squirrel.Eq{
			"email": emails,
		})

	return rawRequest.ToSql()
}

func prepareGetUsersQuery() (string, []interface{}, error) {
	psqlSq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	rawRequest := psqlSq.Select("email", "first_name", "last_name", "password_hash").
		From(tableUsers)

	return rawRequest.ToSql()
}
