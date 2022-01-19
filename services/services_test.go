package services

// Unit test with mock repository layer.
// Indeed, DB is not required: SQLite DB file services/items.db is not created when running unit tests.
// TODO: increase test coverage

import (
	"errors"
	"reflect"
	"testing"

	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/test"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
func TestNewItemService(t *testing.T) {
	tests := []struct {
		name string
		want *ItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemService_CreateItem(t *testing.T) {
	type fields struct {
		r *repository.ItemRepository
	}
	type args struct {
		input models.CreateItemInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				r: tt.fields.r,
			}
			got, err := s.CreateItem(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemService_DeleteItem(t *testing.T) {
	type fields struct {
		r *repository.ItemRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				r: tt.fields.r,
			}
			if err := s.DeleteItem(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_itemService_FindItem(t *testing.T) {
	type fields struct {
		r *repository.ItemRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Item
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				r: tt.fields.r,
			}
			got, got1, err := s.FindItem(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindItem() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_itemService_GetItemRepository(t *testing.T) {
	type fields struct {
		r *repository.ItemRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   repository.ItemRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				r: tt.fields.r,
			}
			if got := s.GetItemRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func Test_itemService_ListItems(t *testing.T) {
	sold := true
	tests := []struct {
		name    string
		want    []models.Item
		wantErr bool
		err     error
	}{
		{"success", []models.Item{models.Item{Sold: &sold}}, false, nil},
		{"error", []models.Item{}, true, errors.New("some error")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := test.NewItemRepository(mocket.DriverName, "connection_string")
			repo.On("ListItems", mock.Anything).Return(tt.want, tt.err)
			s := *NewItemService(repo)

			got, err := s.ListItems()

			assert.Equal(t, tt.want, got)

			if (err != nil) != tt.wantErr {
				t.Errorf("ListItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListItems() got = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func Test_itemService_SetItemRepository(t *testing.T) {
	type fields struct {
		r *repository.ItemRepository
	}
	type args struct {
		r repository.ItemRepository
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				r: tt.fields.r,
			}
		})
	}
}

func Test_itemService_UpdateItem(t *testing.T) {
	type fields struct {
		r *repository.ItemRepository
	}
	type args struct {
		id    string
		input models.UpdateItemInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				r: tt.fields.r,
			}
			got, err := s.UpdateItem(tt.args.id, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
