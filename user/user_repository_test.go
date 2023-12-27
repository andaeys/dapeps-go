package user_test

import (
	"dapeps-go/user"
	"reflect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("User repo test", func() {
	var (
		userRepo user.UserRepositoryImpl
		mockedDB    *MockGormDB
	)

	BeforeSuite(func() {
		mockedDB = new(MockGormDB)
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
				result_users, result_err = userRepo.GetUsers()
				Expect(result_err).NotTo(HaveOccurred())
				Expect(0).To(Equal(len(result_users)))
			})
		})
	})
})

type MockGormDB struct {
    mock.Mock
}

func (m *MockGormDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
    args := m.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (m *MockGormDB) Create(value interface{}) *gorm.DB {
    args := m.Called(value)
    return args.Get(0).(*gorm.DB)
}

func (m *MockGormDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
    args := m.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

