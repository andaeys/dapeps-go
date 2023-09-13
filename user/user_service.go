package user

type UserService interface {
	// CreateUser(name, email string) (*User, error)
	GetUsers() ([]User, error)
	// GetUserByEmail(email string)
}

type UserServiceImpl struct {
	UserRepository UserRepository
}

// func (s *UserServiceImpl) CreateUser(name, email string) (*User, error) {
// 	user := &User{
// 		Name:  name,
// 		Email: email,
// 	}
// 	if err := s.UserRepository.Create(user); err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

func (s *UserServiceImpl) GetUsers() ([]User, error) {
	users, err := s.UserRepository.GetUsers()
	return users, err
}
