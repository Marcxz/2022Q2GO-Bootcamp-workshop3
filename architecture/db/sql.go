package db

import "fmt"


type dbSQL struct {
}

func NewSQL() DB {
	return &dbSQL{}
}

func (db *dbSQL) GetInfoFromDB(id string) (string, error) {
	return fmt.Sprint("Getting info from SQL"), nil
}
