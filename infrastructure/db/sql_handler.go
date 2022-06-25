package db

import (
	"log"
	"os"
	"time"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"github.com/HondaAo/go_blog_app/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLHandler struct {
	db *gorm.DB
}

func NewSqlHandler() gateways.SQLHandler {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	dsn := "host=db user=postgres password=root dbname=development port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err.Error)
	}
	db.AutoMigrate(&entity.Video{}, &entity.Like{}, &entity.User{})
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

func (handler *SQLHandler) Updates(out interface{}) *gorm.DB {
	return handler.db.Updates(out)
}

func (handler *SQLHandler) Delete(out interface{}) *gorm.DB {
	return handler.db.Delete(out)
}
