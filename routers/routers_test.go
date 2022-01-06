package routers
// Unit tests [with setup/teardown] to verify every route, real database is used.
// TODO: fix clearing DB in Test_ListItems_EmptyResult, till then -
// a workaround: delete routers/items.db before running tests.

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/handlers"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/nsavelyeva/go-shopping/repository"
	"github.com/nsavelyeva/go-shopping/test"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")

	var r repository.ItemRepository
	r = *repository.NewItemRepository()
	test.R = r

	// Return a function to teardown the test
	return func(tb testing.TB) {
		log.Println("teardown suite")
		test.R.GetDB().Close()
	}
}

// Almost the same as the above, but this one is for single test instead of collection of tests
func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	return func(tb testing.TB) {
		log.Println("teardown test")
		test.R.ClearTable()
	}
}

func Test_ListItems_EmptyResult(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	req, w := setListItemsRouter(test.DB)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.ItemsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	a.Equal([]models.Item{}, actual.Data)
}

func Test_FindItem_OK(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	a := assert.New(t)

	item, err := insertTestItem(test.DB)
	if err != nil {
		a.Error(err)
	}

	req, w := setFindItemRouter(test.DB,"/items/1")

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.ItemResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.Data.Model = gorm.Model{}
	expected := item
	expected.Model = gorm.Model{}
	a.Equal(expected.Name, actual.Data.Name)
	a.Equal(expected.Price, actual.Data.Price)
	a.NotNil(actual.Data.Sold)
	a.NotNil(actual.Data.ID)
	a.NotNil(actual.Data.CreatedAt)
	a.NotNil(actual.Data.UpdatedAt)
	a.Nil(actual.Data.DeletedAt)
}

func Test_CreateItem_OK(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	a := assert.New(t)
	item := models.CreateItemInput{
		Name: "test",
		Price: 10.99,
	}

	reqBody, err := json.Marshal(item)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateItemRouter(test.DB, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	var actual models.ItemResponse
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Data.Model = gorm.Model{}
	a.Equal(item.Name, actual.Data.Name)
	a.Equal(item.Price, actual.Data.Price)
	a.Equal(false, actual.Data.Sold)
	a.NotNil(actual.Data.ID)
	a.NotNil(actual.Data.CreatedAt)
	a.NotNil(actual.Data.UpdatedAt)
	a.Equal(actual.Data.CreatedAt, actual.Data.UpdatedAt)
	a.Nil(actual.Data.DeletedAt)
}

func setListItemsRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	var h = handlers.NewProvider()
	r.GET("/items", h.ListItems)
	req, err := http.NewRequest(http.MethodGet, "/items", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
}

func setCreateItemRouter(db *gorm.DB,
	body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	var h = handlers.NewProvider()
	r.POST("/items", h.CreateItem)
	req, err := http.NewRequest(http.MethodPost, "/items", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

func setFindItemRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	var h = handlers.NewProvider()
	r.GET("/items/:id", h.FindItem)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
}

func insertTestItem(db *gorm.DB) (*models.Item, error){
	input := models.CreateItemInput{
		Name: "test item",
		Price: 100.0,
	}
	item, err := test.R.CreateItem(&input)
	if err != nil {
		return &models.Item{}, err
	}

	return item, nil
}
