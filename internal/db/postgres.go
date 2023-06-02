package db

import (
	"fmt"
	"github.com/duramash/go-halyklife-tt-2/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "db"
	port     = 6565
	user     = "postgres"
	password = "secret"
	dbname   = "postgres"
)

func GetDB() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := db.AutoMigrate(&types.Author{}, &types.Book{}, &types.Member{}); err != nil {
		fmt.Println(err.Error())
	}
	return db
}
