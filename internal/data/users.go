package data

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	FirstName    string    `db:"first_name"`
	MiddleName   string    `db:"middle_name"`
	LastName     string    `db:"last_name"`
	Birthday     time.Time `db:"birthday"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
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
