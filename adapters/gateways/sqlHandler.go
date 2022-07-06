package gateways

import (
	"gorm.io/gorm"
)

type SQLHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Updates(interface{}, ...interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	WhereCategory(interface{}, ...interface{}) *gorm.DB
	WhereSeries(interface{}, ...interface{}) *gorm.DB
	WhereUserId(interface{}, ...interface{}) *gorm.DB
}
