// /path/to/your/project/user/user_repository_test.go

package user_test

import (
	"dapeps-go/db"
	"dapeps-go/user"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("UserRepositoryImpl", func() {
	var (
		mockDB         *db.MockGormDB
		userRepository user.UserRepository
	)

	BeforeEach(func() {
		mockDB = db.NewMockGormDB()
		userRepository = user.NewUserRepository(mockDB)
	})

	Describe("GetUsers", func() {
		It("should return list of users", func() {
			expectedUsers := []user.User{
				{Name: "John Doe", Email: "john@example.com"},
				{Name: "Jane Doe", Email: "jane@example.com"},
			}

			var datas []interface{}
			for _, u := range expectedUsers {
				datas = append(datas, u)
			}

			er := errors.New("Hello")

			mockDB.SetResult("Find", &gorm.DB{}, datas, er)

			// Call the method under test
			users, err := userRepository.GetUsers()

			// Check expectations
			Expect(err).To(BeNil())
			Expect(expectedUsers).To(ConsistOf(users))
		})
	})
})
