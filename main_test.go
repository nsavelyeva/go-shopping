package main

import (
	"encoding/json"
	"github.com/nsavelyeva/go-shopping/database"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/routers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItemRoute(t *testing.T) {
	database.Setup()
	router := routers.Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)
	router.ServeHTTP(w, req)

	var want models.ItemResponse
	var got models.ItemResponse

	json.Unmarshal([]byte(`{"data":{"name":"Aladdin's lamp","price":999,"sold":true}}`), &want)
	json.Unmarshal([]byte(w.Body.String()), &got)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, want.Data.ID, got.Data.ID)
	assert.Equal(t, want.Data.Name, got.Data.Name)
	assert.Equal(t, want.Data.Price, got.Data.Price)
	assert.Equal(t, want.Data.Sold, got.Data.Sold)
}
