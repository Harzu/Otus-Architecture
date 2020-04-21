package entities

type User struct {
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
}

func (u User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":         u.Email,
		"password_hash": u.PasswordHash,
		"first_name":    u.FirstName,
		"last_name":     u.LastName,
	}
}
