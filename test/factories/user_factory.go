package factories

import (
	"golang-api-server-template/internal/models"
	"golang-api-server-template/test/testutils"
	// "golang.org/x/crypto/bcrypt"
)

// CreateUserFactory generates a user and saves it to the test database
// Email: random, Password: "password"
func UserFactory() models.User {
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	email := testutils.GenerateRandomString(5) + "@example.com"
	user := models.User{
		Name:  testutils.GenerateRandomString(10),
		Tel:   testutils.GenerateRandomStringPtr(11),
		Email: &email,
		// PasswordHash: string(hashedPassword),
	}
	testutils.TestDB.Create(&user)
	testutils.TestDB.First(&user, user.ID)

	return user
}

func UsersFactory(count int) []models.User {
	var users []models.User

	for range count {
		users = append(users, UserFactory())
	}

	return users
}
