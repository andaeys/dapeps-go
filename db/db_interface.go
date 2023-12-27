package db

import "gorm.io/gorm"

type DataBase interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
    Create(value interface{}) *gorm.DB
}
