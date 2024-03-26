package entity

type User struct {
	Id               string `db:"id" json:"id"`
	Username         string `db:"username" json:"username"`
	PasswordHash     string `db:"password_hash" json:"password_hash"`
	BirthdayDate     string `db:"birthday_date" json:"birthday_date"`
	Email            string `db:"email" json:"email"`
	IsEmailConfirmed bool   `db:"is_email_confirmed" json:"is_email_confirmed"`
	Role             string `db:"role" json:"role"`
}
