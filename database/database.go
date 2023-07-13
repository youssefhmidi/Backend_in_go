package database

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase interface {
	Init(location string) error
	CreateTable(Model interface{}) error
	Add(ctx context.Context, Input interface{}) *gorm.DB
	FindOneById(ctx context.Context, VarToAssign interface{}, id uint) *gorm.DB
	FindOneByCol(ctx context.Context, VarToAssign interface{}, Col string, Input string) *gorm.DB
	AppendTo(field string, Model interface{}, Paylod interface{}) error
	FindAll(limit int, RespPayload interface{}) (interface{}, error)
	Preload(ModelName string) *gorm.DB
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

func (db *Database) Add(ctx context.Context, Input interface{}) *gorm.DB {
	return db.Database.WithContext(ctx).Create(Input)
}

func (db *Database) FindOneById(ctx context.Context, VarToAssign interface{}, Id uint) *gorm.DB {
	return db.Database.WithContext(ctx).First(VarToAssign, Id)
}

func (db *Database) FindOneByCol(ctx context.Context, VarToAssign interface{}, Col string, Input string) *gorm.DB {
	return db.Database.WithContext(ctx).First(VarToAssign, fmt.Sprintf("%v = ?", Col), Input)
}

func (db *Database) AppendTo(field string, Model interface{}, Paylod interface{}) error {
	return db.Database.Model(Model).Association(field).Append(Paylod)
}

func (db *Database) FindAll(limit int, RespPayload interface{}) (interface{}, error) {
	reslut := db.Database.Limit(limit).Find(&RespPayload)
	return RespPayload, reslut.Error
}

func (db *Database) Preload(ModelName string) *gorm.DB {
	return db.Database.Preload(ModelName)
}
