package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/kish1n/GiAuth/internal/data"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const usersTable = "users"

type users struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
	counter  squirrel.SelectBuilder
	rank     squirrel.SelectBuilder
}

func NewUsers(db *pgdb.DB) data.UsersQ {
	return &users{
		db:       db,
		selector: squirrel.Select("*").From(usersTable),
		updater:  squirrel.Update(usersTable),
		counter:  squirrel.Select("COUNT(*) as count").From(usersTable),
	}
}

func (q *users) New() data.UsersQ {
	return NewUsers(q.db)
}

func (q *users) Insert(usr data.User) error {
	stmt := squirrel.Insert(usersTable).SetMap(map[string]interface{}{
		"username":      usr.Username,
		"email":         usr.Email,
		"password_hash": usr.PasswordHash,
		"first_name":    usr.FirstName,
		"last_name":     usr.LastName,
		"middle_name":   usr.MiddleName,
	})

	if err := q.db.Exec(stmt); err != nil {
		return fmt.Errorf("insert %s %+v: %w", usersTable, usr, err)
	}

	return nil
}

func (q *users) Update(fields map[string]any) error {
	if err := q.db.Exec(q.updater.SetMap(fields)); err != nil {
		return fmt.Errorf("update %s: %w", usersTable, err)
	}

	return nil
}

func (q *users) Count() (int64, error) {
	res := struct {
		Count int64 `db:"count"`
	}{}

	if err := q.db.Get(&res, q.counter); err != nil {
		return 0, fmt.Errorf("get %s: %w", usersTable, err)
	}

	return res.Count, nil
}

func (q *users) Select() ([]data.User, error) {
	var res []data.User

	if err := q.db.Select(&res, q.selector); err != nil {
		return nil, fmt.Errorf("select balances: %w", err)
	}

	return res, nil
}

func (q *users) Get() (*data.User, error) {
	var res data.User

	if err := q.db.Get(&res, q.selector); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get balance: %w", err)
	}

	return &res, nil
}

func (q *users) FilterByUsername(username ...string) data.UsersQ {
	return q.applyCondition(squirrel.Eq{"username": username})
}

func (q *users) FilterByEmail(email ...string) data.UsersQ {
	return q.applyCondition(squirrel.Eq{"email": email})
}

func (q *users) applyCondition(cond squirrel.Sqlizer) data.UsersQ {
	q.selector = q.selector.Where(cond)
	q.updater = q.updater.Where(cond)
	q.rank = q.rank.Where(cond)
	q.counter = q.counter.Where(cond)
	return q
}
