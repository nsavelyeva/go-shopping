package handlers
// Unit tests to verify handlers functions without involving routers layer, repository and service layers are mock.
// See also https://newbedev.com/go-gin-unit-test-code-example

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)
/*
func TestNewProvider(t *testing.T) {
	tests := []struct {
		name string
		want *Provider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProvider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_CreateItem(t *testing.T) {
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
			p := &Provider{
				s: tt.fields.s,
			}
		})
	}
}

func TestProvider_DeleteItem(t *testing.T) {
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
			p := &Provider{
				s: tt.fields.s,
			}
		})
	}
}

func TestProvider_FindItem(t *testing.T) {
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
			p := &Provider{
				s: tt.fields.s,
			}
		})
	}
}

func TestProvider_GetItemService(t *testing.T) {
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
			p := &Provider{
				s: tt.fields.s,
			}
			if got := p.GetItemService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func TestProvider_ListItems(t *testing.T) {
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

			repo := test.NewItemRepository()
			serv := test.NewItemService(repo)
			serv.On("SetItemRepository", mock.Anything).Return()
			serv.On("ListItems", mock.Anything).Return(tt.data, tt.err)

			p := *NewProvider(serv, repo)
			p.ListItems(c)

			assert.Equal(t, tt.wantCode, w.Code)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}
/*
func TestProvider_SetItemService(t *testing.T) {
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
			p := &Provider{
				s: tt.fields.s,
			}
		})
	}
}

func TestProvider_UpdateItem(t *testing.T) {
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
			p := &Provider{
				s: tt.fields.s,
			}
		})
	}
}
*/
