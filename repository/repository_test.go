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
	"fmt"
	"testing"

	"github.com/nsavelyeva/go-shopping/models"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupMockRepository() *ItemRepository {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	dialect := mysql.New(mysql.Config{
		DSN:                       "mockdb",
		DriverName:                mocket.DriverName,
		SkipInitializeWithVersion: true,
	})

	r := *NewItemRepository(dialect, new(gorm.Config))
	return &r
}

// Tests for ListItems()

func Test_repository_ListItems_Found(t *testing.T) {
	wantReply := []map[string]interface{}{
		{"name": "item 1", "price": float32(10), "sold": true},
		{"name": "item 2", "price": float32(20), "sold": false},
	}
	q := "SELECT * FROM `items` WHERE `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithReply(wantReply)

	r := *SetupMockRepository()
	items, err := r.ListItems()

	assert.Nil(t, err)
	for i := range wantReply {
		assert.Equal(t, wantReply[i]["name"], *items[i].Name)
		assert.Equal(t, wantReply[i]["price"], *items[i].Price)
		assert.Equal(t, wantReply[i]["sold"], *items[i].Sold)

		assert.NotNil(t, items[i].ID)
		assert.NotNil(t, items[i].CreatedAt)
		assert.NotNil(t, items[i].UpdatedAt)
		assert.False(t, items[i].DeletedAt.Valid)
	}
}

func Test_repository_ListItems_NotFound(t *testing.T) {
	q := "SELECT * FROM `items` WHERE `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(0)

	r := *SetupMockRepository()
	items, err := r.ListItems()

	assert.Equal(t, []models.Item{}, items)
	assert.Nil(t, err)
}

func Test_repository_ListItems_ErrorSQL(t *testing.T) {
	wantErr := errors.New("some SQL error")
	q := "SELECT * FROM `items` WHERE `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithError(wantErr)

	r := *SetupMockRepository()
	items, err := r.ListItems()

	assert.Nil(t, items)
	assert.Equal(t, wantErr, err)
}

// Tests for FindItem()

func Test_repository_FindItem_Found(t *testing.T) {
	wantReply := []map[string]interface{}{{"name": "first", "price": 100, "sold": true}}
	q := "SELECT * FROM `items` WHERE `items`.`id` = 1 AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithReply(wantReply)

	r := *SetupMockRepository()
	item, found, err := r.FindItem(1)

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
	item, found, err := r.FindItem(1)

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Nil(t, err)
}

func Test_repository_FindItem_ErrorSQL(t *testing.T) {
	wantErr := errors.New("some SQL error")
	q := "SELECT * FROM `items` WHERE `items`.`id` = 1 AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithError(wantErr)

	r := *SetupMockRepository()
	item, found, err := r.FindItem(1)

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}

func Test_repository_FindItem_ErrorBadIDs(t *testing.T) {
	r := *SetupMockRepository()

	for i := range []int{-1, 0} {
		q := fmt.Sprintf("SELECT * FROM `items` WHERE `items`.`id` = %d AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1", i)
		mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(0)

		item, found, err := r.FindItem(i)

		assert.Nil(t, item)
		assert.False(t, found)
		assert.Nil(t, err)
	}
}

func Test_repository_FindItem_ErrorGo(t *testing.T) {
	wantErr := errors.New("broken item found")
	wantReply := []map[string]interface{}{{"name": "broken", "foo": "bar", "baz": true}}

	q := "SELECT * FROM `items` WHERE `items`.`id` = 1 AND `items`.`deleted_at` IS NULL ORDER BY `items`.`id` LIMIT 1"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithReply(wantReply)

	r := *SetupMockRepository()
	item, found, err := r.FindItem(1)

	assert.False(t, found)
	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}

// Tests for CreateItem()

func Test_repository_CreateItem_Created(t *testing.T) {
	name := "item"
	price := float32(20)
	sold := true
	input := &models.Item{
		Name:  &name,
		Price: &price,
		Sold: &sold,
	}
	q := "INSERT INTO `items` (`created_at`,`updated_at`,`deleted_at`,`name`,`price`,`sold`) VALUES (?,?,?,?,?,?)"
	query := mocket.Catcher.Reset().NewMock().WithQuery(q).WithID(1).WithRowsNum(1)
	wantID := uint(query.LastInsertID)
	wantRowsAffected := int64(1)

	r := *SetupMockRepository()
	item, err := r.CreateItem(input)

	assert.Nil(t, err)

	assert.Equal(t, *input.Name, *item.Name)
	assert.Equal(t, *input.Price, *item.Price)
	assert.False(t, *item.Sold)

	assert.Equal(t, wantID, item.ID)
	assert.Equal(t, wantRowsAffected, query.RowsAffected)

	assert.NotNil(t, item.CreatedAt)
	assert.NotNil(t, item.UpdatedAt)
	assert.Equal(t, item.CreatedAt, item.UpdatedAt)
	assert.False(t, item.DeletedAt.Valid)
}

func Test_repository_CreateItem_ErrorSQL(t *testing.T) {
	name := "item"
	price := float32(20)
	sold := true
	input := &models.Item{
		Name:  &name,
		Price: &price,
		Sold: &sold,
	}
	wantErr := errors.New("some SQL error")
	q := "INSERT INTO `items` (`created_at`,`updated_at`,`deleted_at`,`name`,`price`,`sold`) VALUES (?,?,?,?,?,?)"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithError(wantErr)

	r := *SetupMockRepository()
	item, err := r.CreateItem(input)

	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}

// Tests for UpdateItem()

func Test_repository_UpdateItem_Updated(t *testing.T) {
	name := "new name"
	price := float32(20)
	sold := false
	input := &models.Item{
		Name: &name,
		Price: &price,
		Sold:  &sold,
	}
	q := "UPDATE `items` SET `updated_at`=?,`name`=?,`price`=?,`sold`=? WHERE `items`.`id` = ? AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.NewMock().WithQuery(q).WithRowsNum(1)

	r := *SetupMockRepository()
	item, err := r.UpdateItem(2, input)
	assert.Nil(t, err)

	assert.Equal(t, *input.Name, *item.Name)
	assert.Equal(t, *input.Price, *item.Price)
	assert.Equal(t, *input.Sold, *item.Sold)

    assert.False(t, item.DeletedAt.Valid)
}

func Test_repository_UpdateItem_NotFound(t *testing.T) {
	name := "new name"
	price := float32(20)
	sold := false
	input := &models.Item{
		Name: &name,
		Price: &price,
		Sold:  &sold,
	}
	q := "UPDATE `items` SET `updated_at`=?,`name`=?,`price`=?,`sold`=? WHERE `items`.`id` = ? AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(0)

	r := *SetupMockRepository()
	item, err := r.UpdateItem(2, input)

	wantErr := errors.New("no item found to update")
	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}

func Test_repository_UpdateItem_ErrorBadIDs(t *testing.T) {
	wantErr := errors.New("no item found to update")
	q := "UPDATE `items` SET `updated_at`=?,`name`=?,`price`=?,`sold`=? WHERE `items`.`id` = ? AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(0)

	r := *SetupMockRepository()
	for i := range []int{-1, 0} {
		item, err := r.UpdateItem(i, &models.Item{})

		assert.Nil(t, item)
		assert.Equal(t, wantErr, err)
	}
}

func Test_repository_UpdateItem_ErrorSQL(t *testing.T) {
	name := "new name"
	input := &models.Item{Name: &name}

	q := "UPDATE `items` SET `updated_at`=?,`name`=? WHERE `items`.`id` = ? AND `items`.`deleted_at` IS NULL"
	wantErr := errors.New("some SQL error")
	mocket.Catcher.NewMock().WithQuery(q).WithError(wantErr)

	r := *SetupMockRepository()
	item, err := r.UpdateItem(2, input)
	assert.Nil(t, item)
	assert.Equal(t, wantErr, err)
}

// Tests for DeleteItem()

func Test_repository_DeleteItem_Found(t *testing.T) {
	q := "UPDATE `items` SET `deleted_at`=? WHERE id = ?  AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(1)

	r := *SetupMockRepository()
	err := r.DeleteItem(1)

	assert.Nil(t, err)
}

func Test_repository_DeleteItem_NotFound(t *testing.T) {
	wantErr := errors.New("no item found to delete")
	q := "UPDATE `items` SET `deleted_at`=? WHERE id = ?  AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(0)

	r := *SetupMockRepository()
	err := r.DeleteItem(1)

	assert.Equal(t, wantErr, err)
}

func Test_repository_DeleteItem_ErrorSQL(t *testing.T) {
	wantErr := errors.New("some SQL error")
	q := "UPDATE `items` SET `deleted_at`=? WHERE id = ?  AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithError(wantErr)

	r := *SetupMockRepository()
	err := r.DeleteItem(1)

	assert.Equal(t, wantErr, err)
}

func Test_repository_DeleteItem_ErrorBadIDs(t *testing.T) {
	wantErr := errors.New("no item found to delete")
	q := "UPDATE `items` SET `deleted_at`=? WHERE id = ?  AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(0)

	r := *SetupMockRepository()

	for i := range []int{-1, 0} {
		err := r.DeleteItem(i)
		assert.Equal(t, wantErr, err)
	}
}

func Test_repository_DeleteItem_ErrorGo(t *testing.T) {
	q := "UPDATE `items` SET `deleted_at`=? WHERE id = ?  AND `items`.`deleted_at` IS NULL"
	mocket.Catcher.Reset().NewMock().WithQuery(q).WithRowsNum(1)

	r := *SetupMockRepository()
	err := r.DeleteItem(1)

	assert.Nil(t, err)
}
