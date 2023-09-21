package wrapper

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type DBInterface interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

// Mock

type MockDB struct {
	mock.Mock
}

func NewMockDB() *MockDB {
	return &MockDB{}
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

//Real

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB(d *gorm.DB) *GormDB {
	return &GormDB{DB: d}
}

func (g *GormDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return g.DB.Find(dest, conds...)
}
