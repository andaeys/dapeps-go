package user_test

import (
	"dapeps-go/user"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var _ = Describe("User repo test", func() {
	var (
		userRepo user.UserRepositoryImpl
		mockedDB    *MockDB
	)

	BeforeEach(func() {
		mockedDB = new(MockDB)
		userRepo = user.NewUserRepository(mockedDB)
	})

	//Test Get Users
	Describe("Get Users test", func() {
		var (
			expected_users []user.User
			result_users   []user.User
			result_err     error
			
		)

		Context("when table user has records", func() {

			It("sould return list of users model", func() {
				expected_users = []user.User{
                    {Name: "Momo Oyen", Email: "momo@pepsi.com"},
                    {Name: "Paman mbul", Email: "gembul@pepsi.com"},
                }

				mockedDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{}, nil).Run(func(args mock.Arguments) {
                    dest := args.Get(0)
                    reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(expected_users))
                })

				result_users, result_err = userRepo.GetUsers()
				Expect(result_err).NotTo(HaveOccurred())
				Expect(len(expected_users)).To(Equal(len(result_users)))
			})
		})

		Context("when table user is empty", func() {
			It("sould return empty array", func() {
				expected_users = []user.User{}
				mockedDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{}, nil).Run(func(args mock.Arguments) {
                    dest := args.Get(0)
                    reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(expected_users))
                })
				result_users, result_err = userRepo.GetUsers()
				Expect(result_err).NotTo(HaveOccurred())
				Expect(0).To(Equal(len(result_users)))
			})
		})
	})

	//Test GetByEmail
	Describe("Get by email", func() {
		
		Context("When user is find", func() {
			It("sould return a user data", func ()  {
				expected_user := user.User{Name: "Momo Oyen", Email: "momo@pepsi.com"}
				inputEmail := "momo@pepsi.com"
				mockedDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{}, nil).Run(func(args mock.Arguments){
					dest := args.Get(0)
					reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(expected_user))
				})

				result_user, result_err := userRepo.GetByEmail(inputEmail)

				Expect(result_err).To(BeNil())
				Expect(result_user).To(Equal(result_user))
				mockedDB.AssertExpectations(GinkgoT())
			})
		})

		Context("when a user with the given email does not exist in the database", func() {
			It("should return nil and an error", func() {
				 
				// Arrange
				mockedDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{Error: gorm.ErrRecordNotFound}, nil)

				// Act
				result_user, result_err := userRepo.GetByEmail("nonexistent@example.com")

				// Assert
				Expect(result_err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
				Expect(result_user).To(BeNil())
				mockedDB.AssertExpectations(GinkgoT())
			})
		})
	})
})

type MockDB struct {
    mock.Mock
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
    args := m.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
    args := m.Called(value)
    return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
    args := m.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

