package main
// A quick [smoke] test that verifies a single API route and uses real database (read).
// The test starts and terminates HTTP server automatically, i.e. no need to execute "go run main.go".
// More tests to verify API routes are kept in routers/routers_test.go file.

import (
	"encoding/json"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/routers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItemRoute(t *testing.T) {
	router := routers.Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)
	router.ServeHTTP(w, req)

	var want models.ItemResponse
	var got models.ItemResponse

	json.Unmarshal([]byte(`{"data":{"ID":1,"name":"Aladdin's lamp","price":999,"sold":true}}`), &want)
	json.Unmarshal([]byte(w.Body.String()), &got)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, want.Data.ID, got.Data.ID)
	assert.Equal(t, want.Data.Name, got.Data.Name)
	assert.Equal(t, want.Data.Price, got.Data.Price)
	assert.Equal(t, want.Data.Sold, got.Data.Sold)
}
