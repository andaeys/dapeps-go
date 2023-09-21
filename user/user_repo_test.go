package user

import (
	"dapeps-go/wrapper"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_GetUsers(t *testing.T) {
	// Create a mock DB.
	mockDB := new(wrapper.MockDB)

	// Set up the expectation for Find
	users := []User{{Name: "John Doe"}, {Name: "Jane Doe"}}
	mockDB.On("Find", &users).Return(&gorm.DB{}, nil)

	userRepository := NewUserRepository(mockDB)

	result, err := userRepository.GetUsers()

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
