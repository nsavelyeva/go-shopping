package routers
// TODO: Fix clearing tables, otherwise use workaround:
// remove manually the file routers/items.db before running tests
import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/database"
	"github.com/nsavelyeva/go-shopping/handlers"
	"github.com/nsavelyeva/go-shopping/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ListItems_EmptyResult(t *testing.T) {
	database.Setup()
	db := database.GetDB()
	req, w := setListItemsRouter(db)
	defer db.Close()

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

	a.Equal([]models.Item{}, actual.Items)
	database.ClearTable()
}

func Test_FindItem_OK(t *testing.T) {
	a := assert.New(t)
	database.Setup()
	db := database.GetDB()

	item, err := insertTestItem(db)
	if err != nil {
		a.Error(err)
	}

	req, w := setFindItemRouter(db,"/items/1")
	defer db.Close()

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	//a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.ItemResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	actual.Item.Model = gorm.Model{}
	expected := item
	expected.Model = gorm.Model{}
	a.Equal(expected.Name, actual.Item.Name)
	a.Equal(expected.Price, actual.Item.Price)
	a.NotNil(actual.Item.Sold)
	a.NotNil(actual.Item.ID)
	a.NotNil(actual.Item.CreatedAt)
	a.NotNil(actual.Item.UpdatedAt)
	a.Nil(actual.Item.DeletedAt)
	database.ClearTable()
}

func Test_CreateItem_OK(t *testing.T) {
	a := assert.New(t)
	database.Setup()
	db := database.GetDB()
	item := models.CreateItemInput{
		Name: "test",
		Price: 10.99,
	}

	reqBody, err := json.Marshal(item)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateItemRouter(db, bytes.NewBuffer(reqBody))
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

	actual.Item.Model = gorm.Model{}
	a.Equal(item.Name, actual.Item.Name)
	a.Equal(item.Price, actual.Item.Price)
	a.Equal(false, actual.Item.Sold)
	a.NotNil(actual.Item.ID)
	a.NotNil(actual.Item.CreatedAt)
	a.NotNil(actual.Item.UpdatedAt)
	a.Equal(actual.Item.CreatedAt, actual.Item.UpdatedAt)
	a.Nil(actual.Item.DeletedAt)
	database.ClearTable()
}

func setListItemsRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	api := &handlers.APIEnv{DB: db}
	r.GET("/items", api.ListItems)
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
	api := &handlers.APIEnv{DB: db}
	r.POST("/items", api.CreateItem)
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
	api := &handlers.APIEnv{DB: db}
	r.GET("/items/:id", api.FindItem)

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
	api := &handlers.APIEnv{DB: db}
	input := models.CreateItemInput{
		Name: "test item",
		Price: 100.0,
	}
	item, err := database.CreateItem(api.DB, &input)
	if err != nil {
		return &models.Item{}, err
	}

	return item, nil
}
