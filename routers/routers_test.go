package routers
// Unit tests [with setup/teardown] to verify every route, no real database is used.

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/handlers"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/nsavelyeva/go-shopping/services"
	"github.com/nsavelyeva/go-shopping/test"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	var r repository.ItemRepository
	r = *repository.NewItemRepository(mocket.DriverName, "connection_string")

	test.R = r

	// Return a function to teardown the test
	return func(tb testing.TB) {
		log.Println("teardown suite")
		test.R.GetDB().Close()
	}
}

/*
// Almost the same as the above, but this one is for a single test instead of collection of tests
func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	return func(tb testing.TB) {
		log.Println("teardown test")
		// test.R.ClearTable()
	}
}
*/

func Test_ListItems_EmptyResult(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithRowsNum(0)

	req, w := setListItemsRouter(test.DB)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	var actual models.ItemsResponse
	err = json.Unmarshal(body, &actual)
	assert.Nil(t, err)

	assert.Equal(t, []models.Item{}, actual.Data)
}

func Test_ListItems_NonEmptyResult(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	var expected models.ItemsResponse
	wantBody := `{"data":[{"name":"test-1","price":100.0,"sold":false},{"name":"test-2","price":100.991,"sold":true}]}`
	err := json.Unmarshal([]byte(wantBody), &expected)
	assert.Nil(t, err)
	wantReply := []map[string]interface{}{
		{"name":"test-1","price":100.0,"sold":false},
		{"name":"test-2","price":100.991,"sold":true},
	}
	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithReply(wantReply)

	req, w := setListItemsRouter(test.DB)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	var actual models.ItemsResponse
	err = json.Unmarshal(body, &actual)
	assert.Nil(t, err)

	assert.Equal(t, expected.Data, actual.Data)
}

func Test_FindItem_OK(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	var expected models.Item
	wantBody := `{"name": "first", "price": 100, "sold": true}`
	err := json.Unmarshal([]byte(wantBody), &expected)
	assert.Nil(t, err)

	wantReply := []map[string]interface{}{{"name": "first", "price": 100, "sold": true}}
	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT items.*`).WithReply(wantReply)

	req, w := setFindItemRouter(test.DB,"/items/1")

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	actual := models.ItemResponse{}
	err = json.Unmarshal(body, &actual)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual.Data)
	assert.NotNil(t, actual.Data.CreatedAt)
	assert.NotNil(t, actual.Data.UpdatedAt)
	assert.Nil(t, actual.Data.DeletedAt)
}

func Test_CreateItem_OK(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	item := models.CreateItemInput{
		Name: "test",
		Price: 10.99,
	}

	reqBody, err := json.Marshal(item)
	assert.Nil(t, err)

	req, w, err := setCreateItemRouter(test.DB, bytes.NewBuffer(reqBody))
	assert.Nil(t, err)

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	var actual models.ItemResponse
	err = json.Unmarshal(body, &actual)
	assert.Nil(t, err)

	actual.Data.Model = gorm.Model{}
	assert.Equal(t, item.Name, actual.Data.Name)
	assert.Equal(t, item.Price, actual.Data.Price)
	assert.False(t, actual.Data.Sold)
	assert.NotNil(t, actual.Data.ID)
	assert.NotNil(t, actual.Data.CreatedAt)
	assert.NotNil(t, actual.Data.UpdatedAt)
	assert.Equal(t, actual.Data.CreatedAt, actual.Data.UpdatedAt)
	assert.Nil(t, actual.Data.DeletedAt)
}

func setListItemsRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	g := gin.New()
	var s = services.NewItemService(test.R)
	var h = handlers.NewItemHandler(*s)
	g.GET("/items", h.ListItems)
	req, err := http.NewRequest(http.MethodGet, "/items", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)
	return req, w
}

func setCreateItemRouter(db *gorm.DB,
	body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	g := gin.New()
	var s = services.NewItemService(test.R)
	var h = handlers.NewItemHandler(*s)
	g.POST("/items", h.CreateItem)
	req, err := http.NewRequest(http.MethodPost, "/items", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)
	return req, w, nil
}

func setFindItemRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder) {
	g := gin.New()
	var s = services.NewItemService(test.R)
	var h = handlers.NewItemHandler(*s)
	g.GET("/items/:id", h.FindItem)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)
	return req, w
}
