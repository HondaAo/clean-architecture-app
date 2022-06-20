package gateways

import (
	"gorm.io/gorm"
)

type SQLHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Updates(interface{}) *gorm.DB
}
