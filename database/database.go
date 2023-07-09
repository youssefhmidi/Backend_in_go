package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase interface {
	Init(location string) error
	CreateTable(Model interface{}) error
	AddRow(Input interface{})
	FindOneById(VarToAssign *interface{}, id uint)
	FindOneByCol(VarToAssign interface{}, Col string, Input string) *gorm.DB
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

func (db *Database) Add(Input interface{}) *gorm.DB {
	return db.Database.Create(Input)
}

func (db *Database) FindOneByID(VarToAssign interface{}, Id uint) *gorm.DB {
	return db.Database.First(VarToAssign, Id)
}

func (db *Database) FindOneByCol(VarToAssign interface{}, Col string, Input string) *gorm.DB {
	return db.Database.First(VarToAssign, fmt.Sprintf("%v = ?", Col), Input)
}
