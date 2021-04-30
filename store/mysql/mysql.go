package mysql

import (
	"database/sql"
)

type LinkStore struct {
	db *sql.DB
}

func (m LinkStore) AddLinkMapping(original, shortID string) bool {
	panic("implement me")
}

func (m LinkStore) GetOriginal(shortID string) string {
	panic("implement me")
}

func New() *LinkStore {
	return &LinkStore{db: nil}
}