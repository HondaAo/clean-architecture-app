package db

import (
	"log"
	"os"
	"time"

	"github.com/HondaAo/go_blog_app/adapters/gateways"
	"github.com/HondaAo/go_blog_app/domain/entity"
	"gorm.io/driver/mysql"
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
	dsn := "root:root@tcp(db:3306)/videos-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

func (handler *SQLHandler) Updates(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Where(where).Updates(out)
}

func (handler *SQLHandler) Delete(out interface{}) *gorm.DB {
	return handler.db.Delete(out)
}

func (handler *SQLHandler) WhereCategory(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Where("category = ?", where...).Find(out)
}

func (handler *SQLHandler) WhereSeries(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Where("series = ?", where...).Find(out)
}

func (handler *SQLHandler) WhereUserId(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Where("user_id = ?", where).Find(out)
}
