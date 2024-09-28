package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `db:"id"`
	Username      string    `db:"username"`
	FirstName     string    `db:"first_name"`
	MiddleName    string    `db:"middle_name"`
	LastName      string    `db:"last_name"`
	Email         string    `db:"email"`
	EmailStatus   bool      `db:"email_status"`
	TwoFactorAuth bool      `db:"two_factor_auth"`
	PasswordHash  string    `db:"password_hash"`
	SecretKey     string    `db:"secret_key"`
	Birthday      time.Time `db:"birthday"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type UsersQ interface {
	New() UsersQ

	Insert(q User) error
	Update(map[string]any) error

	Select() ([]User, error)
	Get() (*User, error)
	Count() (int64, error)

	FilterByUsername(...string) UsersQ
	FilterByEmail(...string) UsersQ
}
