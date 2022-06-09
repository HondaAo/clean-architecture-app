package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLHandler struct {
	db *gorm.DB
}

func NewSqlHandler() gateways.SQLHandler {
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// containerName := os.Getenv("DB_CONTAINER")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/development", "root", "root", "db")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
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
