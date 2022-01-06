package test

import (
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/repository"
)

var DB *gorm.DB  // Global variable of a database to be used in test/suite setup/teardown.
var R repository.ItemRepository