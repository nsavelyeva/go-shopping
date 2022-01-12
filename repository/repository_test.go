package repository
// Unit tests to verify SQL queries, no real DB is used thanks to go-mocket.
// See also https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b
// https://github.com/Selvatico/go-mocket/blob/master/DOCUMENTATION.md
//
// In tests, prepare expected replies as follows:
// wantReply := []map[string]interface{}{{"ID": 1, "name": "first", "price": 100, "sold": true}}
// mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithReply(wantReply)
// Important: Use database files here (snake_case) and not struct variables (CamelCase)
// eg: first_name, last_name, date_of_birth NOT FirstName, LastName or DateOfBirth

import (
	"errors"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func SetupRepository() *ItemRepository {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	r := *NewItemRepository(mocket.DriverName, "connection_string")

	return &r
}

func Test_repository_FindItem_Found(t *testing.T) {
	wantReply := []map[string]interface{}{{"ID": 1, "name": "first", "price": 100, "sold": true}}
	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithReply(wantReply)

	r := *SetupRepository()
	item, found, err := r.FindItem("1")

	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "first", item.Name)
	assert.Equal(t, float32(100), item.Price)
	assert.True(t, item.Sold)

	// Now assert the auto-generated fields are available:
	assert.NotNil(t, item.ID)
	assert.NotNil(t, item.CreatedAt)
	assert.NotNil(t, item.UpdatedAt)
	assert.Equal(t, item.CreatedAt, item.UpdatedAt)
	assert.Nil(t, item.DeletedAt)
}

func Test_repository_FindItem_NotFound(t *testing.T) {
	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithRowsNum(0)

	r := *SetupRepository()
	item, found, err := r.FindItem("1")

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Nil(t, err)
}


func Test_repository_FindItem_Error(t *testing.T) {
	wantErr := errors.New("some SQL error")
	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithError(wantErr)

	r := *SetupRepository()
	item, found, err := r.FindItem("1")

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}
