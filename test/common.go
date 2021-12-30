package test

import (
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/database"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuiteEnv struct {
	suite.Suite
	db *gorm.DB
}

// Tests are run before they start
func (suite *TestSuiteEnv) SetupSuite() {
	database.Setup()
	suite.db = database.GetDB()
}

// Running after each test
func (suite *TestSuiteEnv) TearDownTest() {
	database.ClearTable()
}

// Running after all tests are completed
func (suite *TestSuiteEnv) TearDownSuite() {
	suite.db.Close()
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(TestSuiteEnv))
}

