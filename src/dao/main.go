package dao

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type DynamicDao interface {
	Insert(sql string) error
	DeleteTable(tableName string) error
	CreateTable(sql string) error
	ShowTables()
	HasTable(tableName string) bool
	TruncateTable(tableName string) error
}

type DynamicDaoImpl struct {
	DynamicDao
	db *gorm.DB
}

func NewDynamicDaoImpl(db *gorm.DB) DynamicDao {
	dao := DynamicDaoImpl{
		db: db,
	}
	return &dao
}

func (dao *DynamicDaoImpl) Insert(sql string) error {
	db := dao.db.Exec(sql)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (dao *DynamicDaoImpl) ShowTables() {
	dao.db.Raw("show tables")
}

func (dao *DynamicDaoImpl) DeleteTable(tableName string) error {
	db := dao.db.DropTable(tableName)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (dao *DynamicDaoImpl) CreateTable(sql string) error {
	db := dao.db.Exec(sql)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (dao *DynamicDaoImpl) HasTable(tableName string) bool {
	return dao.db.HasTable(tableName)
}

func (dao *DynamicDaoImpl) TruncateTable(tableName string) error {
	db:=dao.db.Exec(fmt.Sprintf("truncate table %s",tableName))
	if db.Error != nil {
		return db.Error
	}
	return nil
}
