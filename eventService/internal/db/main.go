package db

import (
	"eventService/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB
var cfg *configs.Config
var connect string

func init() {
	cfg = configs.Get()
	connect = "host=" + cfg.HostDb +
		" port=" + cfg.PortDb +
		" user= " + cfg.User +
		" password=" + cfg.Password +
		" dbname=" + cfg.DbName +
		" sslmode=" + cfg.SslMode
}

func DB() *gorm.DB {
	return database
}

func Connect() {
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	database = db
	log.Println("Connected to the database")
}
