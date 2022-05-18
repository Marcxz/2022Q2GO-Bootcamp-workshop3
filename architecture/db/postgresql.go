package db

import (
	"fmt"
)

type dbPostgreSQL struct {
}

func NewPostgreSQL() DB {
	return &dbPostgreSQL{}
}
func (db *dbPostgreSQL) GetInfoFromDB(id string) (string, error) {
	return fmt.Sprint("Getting info from PostgreSQL"), nil
}
