package db

import (
	"fmt"
	"os"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLHandler struct {
	db *gorm.DB
}

func NewSqlHandler() gateways.SQLHandler {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, pass, dbName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Find(out, where...)
}

func (handler *SQLHandler) Create(out interface{}) *gorm.DB {
	return handler.db.Create(out)
}
