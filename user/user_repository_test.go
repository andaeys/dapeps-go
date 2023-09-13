package user_test

import (
	"dapeps-go/db"
	"dapeps-go/user"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("User repo test", func() {
	var (
		userRepo user.UserRepositoryImpl
		dbCon    *gorm.DB
	)

	BeforeSuite(func() {
		db.InitDB()
		dbCon = db.GetDB()
		userRepo = user.NewUserRepository(dbCon)
	})

	AfterSuite(func() {
		dbCon.Exec("DELETE FROM users")
	})

	//Test Get Users
	Describe("Get Users test", func() {
		var (
			expected_users []user.User
			result_users   []user.User
			result_err     error
			err            error
		)

		Context("when table user has records", func() {

			BeforeEach(func() {
				// Create multiple users in the database for testing
				user1 := user.User{Name: "Momo Oyen", Email: "momo@pepsi.com"}
				user2 := user.User{Name: "Paman mbul", Email: "gembul@pepsi.com"}
				expected_users[0] = user1
				expected_users[1] = user2

				err = dbCon.Create(&user1).Error
				Expect(err).NotTo(HaveOccurred())

				err = dbCon.Create(&user2).Error
				Expect(err).NotTo(HaveOccurred())
			})

			It("sould return list of users model", func() {
				result_users, result_err = userRepo.GetUsers()
				Expect(result_err).NotTo(HaveOccurred())
				Expect(len(expected_users)).To(Equal(len(result_users)))
			})
		})

		Context("when table user is empty", func() {
			It("sould return empty array", func() {
				result_users, result_err = userRepo.GetUsers()
				Expect(result_err).NotTo(HaveOccurred())
				Expect(0).To(Equal(len(result_users)))
			})
		})
	})
})
