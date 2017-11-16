package local

import (
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/emrekeskinmac/mypass/account"
	_ "github.com/mattn/go-sqlite3"
)

type LocalProvider struct {
	username string
	secret   []byte
	db       *sql.DB
}

func New(username, password, dbPath string) (*LocalProvider, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	db.Exec("create table accounts(username varchar(255), name varchar(255), hex text, createdAt int)")

	return &LocalProvider{
		username: username,
		db:       db,
		secret:   account.MakeSecret(username, password),
	}, nil
}

func (lp *LocalProvider) Save(username, name, hx string) (err error) {
	sql, args, err := squirrel.
		Insert("accounts").
		Columns("username", "name", "hex", "createdAt").
		Values(username, name, hx, time.Now().Unix()).
		ToSql()
	if err != nil {
		return err
	}

	_, err = lp.db.Exec(sql, args...)
	return err
}

func (lp *LocalProvider) Find(username, name string) (hx []string, err error) {
	sql, args, err := squirrel.Select("hex").
		From("accounts").
		Where(squirrel.Eq{"username": username, "name": name}).
		ToSql()

	if err != nil {
		return hx, err
	}

	rows, err := lp.db.Query(sql, args...)
	if err != nil {
		return hx, err
	}

	for rows.Next() {
		var h string
		if err := rows.Scan(&h); err != nil {
			return hx, err
		}
		hx = append(hx, h)
	}
	return hx, rows.Err()
}

func (lp *LocalProvider) Delete(username string, name string) (hx []string, err error) {
	return hx, err
}
