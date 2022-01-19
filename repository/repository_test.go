package repository

// Unit tests to verify SQL queries, no real DB is used thanks to go-mocket.
// See also https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b
// https://github.com/Selvatico/go-mocket/blob/master/DOCUMENTATION.md
//
// In tests, prepare expected replies as follows:
// wantReply := []map[string]interface{}{{"name": "first", "price": 100, "sold": true}}
// mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithReply(wantReply)
// Important: Use database files here (snake_case) and not struct variables (CamelCase)
// eg: first_name, last_name, date_of_birth NOT FirstName, LastName or DateOfBirth

import (
	"errors"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func SetupMockRepository() *ItemRepository {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	dialect := mysql.New(mysql.Config{
		DSN:                             "mockdb",
		DriverName:                      mocket.DriverName,
		SkipInitializeWithVersion: true,
	})

	r := *NewItemRepository(dialect, new(gorm.Config))
	return &r
}

func Test_repository_FindItem_Found(t *testing.T) {
	wantReply := []map[string]interface{}{{"name": "first", "price": 100, "sold": true}}
	q := "SELECT * FROM `items` WHERE `items`.`id` = 1 AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithReply(wantReply)

	r := *SetupMockRepository()
	item, found, err := r.FindItem("1")

	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "first", *item.Name)
	assert.Equal(t, float32(100), *item.Price)
	assert.True(t, *item.Sold)

	assert.NotNil(t, item.ID)
	assert.NotNil(t, item.CreatedAt)
	assert.NotNil(t, item.UpdatedAt)
	assert.Equal(t, item.CreatedAt, item.UpdatedAt)
	assert.False(t, item.DeletedAt.Valid)
}

func Test_repository_FindItem_NotFound(t *testing.T) {
	q := "SELECT * FROM `items` WHERE `items`.`id` = 1 AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(0)

	r := *SetupMockRepository()
	item, found, err := r.FindItem("1")

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Nil(t, err)
}

func Test_repository_FindItem_ErrorSQL(t *testing.T) {
	wantErr := errors.New("some SQL error")
	q := "SELECT * FROM `items` WHERE `items`.`id` = 1 AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithError(wantErr)

	r := *SetupMockRepository()
	item, found, err := r.FindItem("1")

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}

func Test_repository_FindItem_ErrorGo(t *testing.T) {
	wantErr := errors.New("broken item found")
	wantReply := []map[string]interface{}{{"name": "broken", "foo": "bar", "baz": true}}

	q := "SELECT * FROM `items` WHERE `items`.`id` = 1 AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithReply(wantReply)

	r := *SetupMockRepository()
	item, found, err := r.FindItem("1")

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}
