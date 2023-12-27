package db

import "gorm.io/gorm"

type DataBase interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}
