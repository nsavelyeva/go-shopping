package test

import (
	"gorm.io/gorm"
)

// DB is declared to be a global variable that represents a database, it needs access in test setup/teardown functions
var DB *gorm.DB

// R is declared to be a global variable that represents a repository, it needs access in test setup/teardown functions
var R *ItemRepository
