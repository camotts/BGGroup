package store

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	errs "github.com/pkg/errors"
)

type Game struct {
	Model
	Name        string
	Description string
	Image       string
	Thumbnail   string
	Year        int
}

func (s *Store) CreateGame(tx *sqlx.Tx, g *Game) error {
	stmt, err := tx.Prepare("insert into games (id, name, description, image, thumbnail, year) values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return errs.Wrap(err, "unable to perp games insert statement")
	}
	_, err = stmt.Exec(g.Id, g.Name, g.Description, g.Image, g.Thumbnail, g.Year)
	return errs.Wrap(err, "unable to insert game")
}

func (s *Store) CreateGameMulti(tx *sqlx.Tx, gs ...*Game) error {
	var retErr error
	for i := 0; i < len(gs); i += 50 {
		var sec []*Game
		if i > len(gs)-50 {
			sec = gs[i:]
		} else {
			sec = gs[i : i+50]
		}
		valueStrs := make([]string, 0, len(sec))
		valueArgs := make([]interface{}, 0, len(sec))
		for _, g := range sec {
			valueStrs = append(valueStrs, "(?, ?, ?, ?, ?, ?)")
			valueArgs = append(valueArgs, g.Id, g.Name, g.Description, g.Image, g.Thumbnail, g.Year)
		}
		stmt := fmt.Sprintf("insert into games (id, name, description, image, thumbnail, year) values %s", strings.Join(valueStrs, ","))
		if _, err := tx.Exec(stmt, valueArgs); err != nil {
			retErr = errors.Join(retErr, err)
		}

	}
	return retErr
}
