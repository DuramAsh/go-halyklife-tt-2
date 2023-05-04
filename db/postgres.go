package db

import (
	"fmt"
	"github.com/duramash/go-halyklife-tt-2/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwerty123"
	dbname   = "postgres"
)

func GetDB() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("no connection to db")
	}
	_ = db.AutoMigrate(&types.Author{})
	_ = db.AutoMigrate(&types.Book{})
	_ = db.AutoMigrate(&types.Member{})
	return db
}
