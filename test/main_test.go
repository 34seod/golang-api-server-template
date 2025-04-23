package test

import (
	"os"
	"testing"

	"golang-api-server-template/test/testutils"
)

func TestMain(m *testing.M) {
	testutils.SetupTestDatabase()

	// Run all tests
	code := m.Run()

	// Teardown the database connection or other cleanup tasks
	testutils.TearDownTestDatabase()

	// Exit with the test run status code
	os.Exit(code)
}
