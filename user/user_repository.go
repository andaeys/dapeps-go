package user

import "gorm.io/gorm"

type UserRepository interface {
	GetUsers() ([]User, error)
	// GetByEmail(email string) (*User, error)
	// Create(user *User) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryImpl {
	return UserRepositoryImpl{DB: db}
}

// func (r *UserRepositoryImpl) GetByEmail(email string) (*User, error) {
// 	var user User
// 	result := r.DB.First(&user, email)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &user, nil
// }

// func (r *UserRepositoryImpl) Create(user *User) error {
// 	result := r.DB.Create(user)
// 	return result.Error
// }

func (r *UserRepositoryImpl) GetUsers() ([]User, error) {
	var users []User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
