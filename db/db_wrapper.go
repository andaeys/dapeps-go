package db

import (
	"time"

	"gorm.io/gorm"
)

// Declare interface for gorm db function used

type GormDB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
}

// ========================== Mock implementaion for testing purpose ===================================

type MockGormDB struct {
	results map[string]mockResult
}

type mockResult struct {
	result *gorm.DB
	err    error
	users  *[]interface{}
}

func NewMockGormDB() *MockGormDB {
	return &MockGormDB{
		results: make(map[string]mockResult),
	}
}

func (m *MockGormDB) SetResult(method string, result *gorm.DB, datas *[]interface{}, err error) {
	m.results[method] = mockResult{
		result: &gorm.DB{},
		err:    err,
		users:  datas,
	}
}

func (m *MockGormDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	result, ok := m.results["Find"]
	if !ok {
		return &gorm.DB{} // Return a default value
	}
	dest = result.users
	time.Sleep(1000 * time.Millisecond)
	return result.result
}

func (m *MockGormDB) Create(value interface{}) *gorm.DB {
	result, ok := m.results["Create"]
	if !ok {
		return &gorm.DB{} // Return a default value
	}
	return result.result
}

func (m *MockGormDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	result, ok := m.results["First"]
	if !ok {
		return &gorm.DB{} // Return a default value
	}
	return result.result
}

// ========================== Real implementaion using gorm library ===================================

type RealDB struct {
	DB *gorm.DB
}

func NewRealDB(d *gorm.DB) *RealDB {
	return &RealDB{DB: d}
}

func (m *RealDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return m.DB.Find(dest, conds...)
}

func (m *RealDB) Create(value interface{}) *gorm.DB {
	return m.DB.Create(value)
}

func (m *RealDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return m.DB.First(dest, conds...)
}
