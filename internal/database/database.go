package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type DataManager struct {
	connection *sqlx.DB
}

func (db *DataManager) Connect(dbConnString string) {
	var err error
	db.connection, err = sqlx.Connect("mysql", dbConnString)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
}

func NewDataManager() *DataManager {
	return &DataManager{}
}

func (db *DataManager) Close() {
	_ = db.connection.Close()
}
