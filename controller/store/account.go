package store

import (
	"github.com/jmoiron/sqlx"
	errs "github.com/pkg/errors"
)

type Account struct {
	Model
	Email      string
	Password   string
	BGGAccount string
}

func (s *Store) CreateAccount(tx *sqlx.Tx, a *Account) error {
	stmt, err := tx.Prepare("insert into accounts (id, email, password, bgg_account) values ($1, $2, $3, $4)")
	if err != nil {
		return errs.Wrap(err, "unable to perp accounts insert statement")
	}
	_, err = stmt.Exec(a.Id, a.Email, a.Password, a.BGGAccount)
	return errs.Wrap(err, "unable to insert account")
}
