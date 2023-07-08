package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase interface {
	Init(location string) error
	CreateTable(Model interface{}) error
	AddRow(Input interface{})
}
type Database struct {
	Database *gorm.DB
}

func (db *Database) Init(location string) error {
	DB, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	db.Database = DB
	return err
}

func (db *Database) CreateTable(Model interface{}) error {
	err := db.Database.AutoMigrate(&Model)
	return err
}

func (db *Database) AddRow(Input interface{}) {
	db.Database.Create(Input)
}
