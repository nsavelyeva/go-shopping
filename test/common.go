package test

import (
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable of a database to be used in test/suite setup/teardown.
var R *ItemRepository
