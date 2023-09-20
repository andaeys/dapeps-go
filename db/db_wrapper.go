package db

import (
	"gorm.io/gorm"
)

type GormDB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
}

type MockGormDB struct {
	results map[string]mockResult
}

type mockResult struct {
	result *gorm.DB
	err    error
	users  []interface{}
}

func NewMockGormDB() *MockGormDB {
	return &MockGormDB{
		results: make(map[string]mockResult),
	}
}

func (m *MockGormDB) SetResult(method string, result *gorm.DB, datas []interface{}, err error) {
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
