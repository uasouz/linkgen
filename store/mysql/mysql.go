package mysql

import (
	"database/sql"
	"linkgen/store"
	"linkgen/store/mysql/migrations"

	"github.com/pressly/goose/v3"

	_ "github.com/go-sql-driver/mysql"
)

type LinkStore struct {
	db *sql.DB
}

func (m LinkStore) AddLinkMapping(original, shortID string) error {
	_, err := m.db.Query("INSERT INTO links (shortid,originalURL) values(?,?)", shortID, original)
	return err
}

func (m LinkStore) GetOriginal(shortID string) (originalURL string, err error) {
	row := m.db.QueryRow("SELECT originalURL from links where shortid=?", shortID)
	err = row.Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			err = store.ErrLinkNotFound
		}
	}
	return
}

func New(dsn string) (*LinkStore, error) {
	dataBase, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = dataBase.Ping()
	if err != nil {
		return nil, err
	}

	// Migrate the DB
	goose.SetBaseFS(migrations.LinkGenStoreMySQLMigrations)

	if err = goose.SetDialect("mysql"); err != nil {
		return nil, err
	}

	if err = goose.Up(dataBase, "."); err != nil {
		return nil, err
	}

	return &LinkStore{
		db: dataBase,
	}, nil
}
