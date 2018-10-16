package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
	"log"
	"github.com/go-errors/errors"
)

type DB_RESOURCE string

const (
	DB_RESOURCE_LOCAL DB_RESOURCE = "local"
)

var DBPool Resource

type Resource interface {
	GetDB(name DB_RESOURCE) (*gorm.DB, bool)
	PutDB(name DB_RESOURCE, db *gorm.DB) bool
	Size() int
}

type DBResource struct {
	pool map[DB_RESOURCE]*gorm.DB
}

func NewDBResource() Resource {
	resource := new(DBResource)
	resource.pool = make(map[DB_RESOURCE]*gorm.DB)
	return resource
}

func (resource *DBResource) Size() int {
	return len(resource.pool)
}

func (resource *DBResource) GetDB(name DB_RESOURCE) (*gorm.DB, bool) {
	if db, ok := resource.pool[name]; !ok {
		return nil, false
	} else {
		return db, true
	}
}

func (resource *DBResource) PutDB(name DB_RESOURCE, db *gorm.DB) bool {
	if db == nil {
		return false
	}
	if _, ok := resource.pool[name]; ok {
		return false
	} else {
		resource.pool[name] = db
		return true
	}
}

func init() {
	resource := NewDBResource()
	fmt.Println("========连接数据库=========")
	for _, dbConfig := range ConfigurationContext.DBConfigs {
		if !dbConfig.Config.Switch {
			log.Printf("resource [%s] not load", dbConfig.Name)
			continue
		}
		db, err := gorm.Open(dbConfig.Config.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&allowNativePasswords=true", dbConfig.Config.UserName, dbConfig.Config.Password, dbConfig.Config.Host, dbConfig.Config.Port, dbConfig.Config.DataBaseName))
		if err != nil {
			panic(err)
		}
		db.LogMode(dbConfig.Config.Mode)
		db.DB().SetConnMaxLifetime(time.Minute * dbConfig.Config.ConnMaxLifetime)
		db.DB().SetMaxOpenConns(dbConfig.Config.MaxOpenNum)
		db.DB().SetMaxIdleConns(dbConfig.Config.MaxIdleNum)
		ok := resource.PutDB(DB_RESOURCE(dbConfig.Name), db)
		if !ok {
			log.Fatal("the db add failed ...")
		}
	}
	fmt.Printf("数据库初始化完毕,数据源个数:%d\n", resource.Size())
	DBPool = resource
}

func GetDefaultDB() (*gorm.DB, error) {
	if db, ok := DBPool.GetDB(DB_RESOURCE_LOCAL); !ok {
		return nil, errors.New("获取数据源失败")
	} else {
		return db, nil
	}
}
