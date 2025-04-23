package factories

import (
	"golang-api-server-template/internal/model"
	"golang-api-server-template/test/testutils"
	// "golang.org/x/crypto/bcrypt"
)

// CreateUserFactory generates a user and saves it to the test database
// Email: random, Password: "password"
func UserFactory() model.User {
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	email := testutils.GenerateRandomString(5) + "@example.com"
	user := model.User{
		Name:  testutils.GenerateRandomString(10),
		Tel:   testutils.GenerateRandomStringPtr(11),
		Email: &email,
		// PasswordHash: string(hashedPassword),
	}
	testutils.TestDB.Create(&user)
	testutils.TestDB.First(&user, user.ID)

	return user
}

func UsersFactory(count int) []model.User {
	var users []model.User

	for range count {
		users = append(users, UserFactory())
	}

	return users
}
