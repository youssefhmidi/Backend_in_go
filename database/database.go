package database

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SqliteDatabase interface {
	Init(location string) error
	CreateTable(Model interface{}) error
	Add(ctx context.Context, Input interface{}) *gorm.DB
	FindOneById(ctx context.Context, VarToAssign interface{}, id uint) *gorm.DB
	FindOneByCol(ctx context.Context, VarToAssign interface{}, Col string, Input string) *gorm.DB
	AppendTo(field string, Model interface{}, Paylod interface{}, ctx context.Context) error
	FindAll(limit int, RespPayload interface{}) (interface{}, error)
	FindAllByCol(limit int, RespPayload interface{}, Col string, Input any, ctx context.Context) (interface{}, error)
	UpdateRow(ctx context.Context, Model interface{}, Col string, NewVal interface{}) *gorm.DB
	DeleteRecordById(ctx context.Context, Model interface{}, Id uint) *gorm.DB
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
	return db.Database.Preload(clause.Associations).WithContext(ctx).First(VarToAssign, Id)
}

func (db *Database) FindOneByCol(ctx context.Context, VarToAssign interface{}, Col string, Input string) *gorm.DB {
	return db.Database.Preload(clause.Associations).WithContext(ctx).First(VarToAssign, fmt.Sprintf("%v = ?", Col), Input)
}

func (db *Database) AppendTo(field string, Model interface{}, Paylod interface{}, ctx context.Context) error {
	return db.Database.Model(Model).WithContext(ctx).Association(field).Append(Paylod)
}

func (db *Database) FindAll(limit int, RespPayload interface{}) (interface{}, error) {
	reslut := db.Database.Preload(clause.Associations).Limit(limit).Find(&RespPayload)
	return RespPayload, reslut.Error
}

func (db *Database) FindAllByCol(limit int, RespPayload interface{}, Col string, Input any, ctx context.Context) (interface{}, error) {
	reslut := db.Database.Preload(clause.Associations).WithContext(ctx).Limit(limit).Find(&RespPayload, fmt.Sprintf("%v = ?", Col), Input)
	return RespPayload, reslut.Error
}
func (db *Database) UpdateRow(ctx context.Context, Model interface{}, Col string, NewVal interface{}) *gorm.DB {
	return db.Database.Model(Model).WithContext(ctx).Update(Col, NewVal)
}

func (db *Database) DeleteRecordById(ctx context.Context, Model interface{}, Id uint) *gorm.DB {
	return db.Database.WithContext(ctx).Delete(Model, Id)
}
