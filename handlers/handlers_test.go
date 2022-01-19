package handlers

// Unit tests to verify handlers functions without involving routers layer, repository and service layers are mock.
// See also https://newbedev.com/go-gin-unit-test-code-example
// TODO: increase test coverage

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/test"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
func TestNewHandler(t *testing.T) {
	tests := []struct {
		name string
		want *itemHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_CreateItem(t *testing.T) {
	type fields struct {
		s *services.ItemService
	}
	type args struct {
		c *gin.Context
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
			p := &itemHandler{
				s: tt.fields.s,
			}
		})
	}
}

func TestHandler_DeleteItem(t *testing.T) {
	type fields struct {
		s *services.ItemService
	}
	type args struct {
		c *gin.Context
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
			p := &itemHandler{
				s: tt.fields.s,
			}
		})
	}
}

func TestHandler_FindItem(t *testing.T) {
	type fields struct {
		s *services.ItemService
	}
	type args struct {
		c *gin.Context
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
			p := &itemHandler{
				s: tt.fields.s,
			}
		})
	}
}

func TestHandler_GetItemService(t *testing.T) {
	type fields struct {
		s *services.ItemService
	}
	tests := []struct {
		name   string
		fields fields
		want   services.ItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &itemHandler{
				s: tt.fields.s,
			}
			if got := p.GetItemService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func TestHandler_ListItems(t *testing.T) {
	tests := []struct {
		name       string
		data       []models.Item
		err        error
		wantCode   int
		wantBody   string
	}{
		{"success", []models.Item{}, nil, 200, `{"data":[]}`},
		{"error", nil, errors.New("some error"), 400, `{"error":"some error"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			repo := test.NewItemRepository(mocket.DriverName, "connection_string")
			repo.On("ListItems", mock.Anything).Return(tt.data, tt.err)
			serv := test.NewItemService(repo)
			serv.On("SetItemRepository", mock.Anything).Return()
			serv.On("ListItems", mock.Anything).Return(tt.data, tt.err)

			p := NewItemHandler(serv)
			p.ListItems(c)

			assert.Equal(t, tt.wantCode, w.Code)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}

/*
func TestHandler_SetItemService(t *testing.T) {
	type fields struct {
		s *services.ItemService
	}
	type args struct {
		s services.ItemService
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
			p := &itemHandler{
				s: tt.fields.s,
			}
		})
	}
}

func TestHandler_UpdateItem(t *testing.T) {
	type fields struct {
		s *services.ItemService
	}
	type args struct {
		c *gin.Context
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
			p := &itemHandler{
				s: tt.fields.s,
			}
		})
	}
}
*/
