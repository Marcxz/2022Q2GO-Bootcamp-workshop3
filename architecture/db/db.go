package db

type DB interface {
	GetInfoFromDB(id string) (string, error)
}
