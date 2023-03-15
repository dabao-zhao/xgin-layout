package data

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/dabao-zhao/xgin-layout/config"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewGreeterRepo)

type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) *Data {
	return &Data{
		db: db,
	}
}

func NewDB(conf *config.Config) *gorm.DB {
	//db, err := gorm.Open(mysql.Open(conf.Db.Dsn()), &gorm.Config{})
	//if err != nil {
	//	log.Fatalf("failed opening connection to mysql: %v", err)
	//}
	//return db
	return nil
}
