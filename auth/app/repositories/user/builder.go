package user

import (
	"otus-auth/app/entities"

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

func prepareUpdateUserQuery(userEmail string, newUserInfo entities.User) (string, []interface{}, error) {
	psqlSq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	rawRequest := psqlSq.Update(tableUsers).
		SetMap(newUserInfo.ToMap()).
		Where(squirrel.Eq{
			"email": userEmail,
		})

	return rawRequest.ToSql()
}

func prepareGetUserQuery(email string) (string, []interface{}, error) {
	psqlSq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	rawRequest := psqlSq.Select("email", "first_name", "last_name", "password_hash").
		From(tableUsers).
		Where(squirrel.Eq{
			"email": email,
		})

	return rawRequest.ToSql()
}
