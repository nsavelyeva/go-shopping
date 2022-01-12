package repository
// The repository layer is responsible for connecting directly to the database to retrieve and/or modify records.

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	//"github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/jinzhu/gorm/dialects/postgres"
	//"github.com/jinzhu/gorm/dialects/sqlite"
	// "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/models"
)

type ItemRepository interface {
	ListItems() ([]models.Item, error)
	FindItem(id string) (*models.Item, bool, error)
	CreateItem(input *models.CreateItemInput) (*models.Item, error)
	UpdateItem(id string, input *models.UpdateItemInput) (*models.Item, error)
	DeleteItem(id string) error
	GetDB() *gorm.DB
	ClearTable()  *gorm.DB
}

type itemRepository struct {
	dn string   // driverName
	cs string   // connectionString
	db *gorm.DB
}

func ConnectDB(driverName string, connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(driverName, connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to the database due to error: %s", err)
	}
	return db, err
}

func NewItemRepository(driverName string, connectionString string) *ItemRepository {
	db, _ := ConnectDB(driverName, connectionString)
	db.LogMode(false)
	db.AutoMigrate(&models.Item{})

	var r ItemRepository = &itemRepository{dn: driverName, cs: connectionString, db: db}
	return &r
}

func (r *itemRepository) SetDB(db gorm.DB) {
	r.db = &db
}

func (r *itemRepository) GetDB() *gorm.DB {
	if r.db.DB().Ping() != nil {
		return r.db
	}
	db, _ := ConnectDB(r.dn, r.cs)
	return db
}

// TODO: fix queries to clear the records in the database
func (r *itemRepository) ClearTable()  *gorm.DB {
	//DB.Lock()
	r.db.Begin()
	r.db.Exec("DELETE FROM `items` WHERE 1=1")
	r.db.Exec("ALTER SEQUENCE items_id_seq RESTART WITH 1")
	//DB.Exec("UPDATE `sqlite_sequence` SET `seq` = 0 WHERE `name` = 'items'")
	r.db.Commit()
	//DB.Unlock()
	//DB.Exec("ALTER SEQUENCE items_id_seq RESTART WITH 1")
	return r.db
}

/*
func NewRepository(r repository, dialect Dialect) (Repository, error) {
	db := dialect.DialectSQLite3("items.db")
	//db, err := gorm.Open("sqlite3", "items.db")

	if err != nil {
		log.Fatal("Failed to connect to the database!")
	}


	return &repository{db}, nil
}

type Dialect func(Repository) Repository

func DialectSQLite3(dbFileName string) Dialect {
	return func(r Repository) Repository {
		db, err := gorm.Open("sqlite3", dbFileName)
		if err != nil {
			log.Fatalf("Cannot connect to SQLite3 using file name '%s' due to %s", dbFileName, err)
		}
		return &repository{db}
	}
}

func DialectMySQL(host string, port string, user string, password string, database string) Dialect {
	return func(r Repository) Repository {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, port, database)
		// e.g. "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"

		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Cannot connect to MySQL using DSN '%s' due to %s", dsn, err)
		}
		return &repository{db}
	}
}

func DialectPostgreSQL(host string, port string, user string, password string, database string) Dialect {
	return func(r Repository) Repository {
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			host, port, user, database, password)
		// e.g. "host=myhost port=myport user=gorm dbname=gorm password=mypassword"

		db, err := gorm.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Cannot connect to PostgreSQL using DSN '%s' due to %s", dsn, err)
		}
		return &repository{db}
	}
}
*/

// Close attaches the provider and close the connection
func (r *itemRepository) Close() {
	r.db.Close()
}

func (r *itemRepository) ListItems() ([]models.Item, error) {
	var items []models.Item
	query := r.GetDB().Select("items.*").
		Group("items.id").
		Find(&items)
	defer query.Close()

	if query.RecordNotFound() {
		return nil, nil
	}
	if err := query.Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *itemRepository) FindItem(id string) (*models.Item, bool, error) {
	var item models.Item
	query := r.GetDB().Select("items.*").
		Group("items.id").
		Where("items.id = ?", id).
		First(&item)
	defer query.Close()

	if query.RecordNotFound() {
		return nil, false, nil
	}
	if err := query.Error; err != nil {
		return nil, false, err
	}

	return &item, true, nil
}

func (r *itemRepository) CreateItem(input *models.CreateItemInput) (*models.Item, error) {
	item := models.Item{
		Name: input.Name,
		Price: input.Price,
		Sold: false,
	}
	if err := r.GetDB().Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) UpdateItem(id string, input *models.UpdateItemInput) (*models.Item, error) {
	item, found, err := r.FindItem(id)
	if err != nil || !found {
		return nil, errors.New("item not found")
	}
	data := models.Item{
		Name: input.Name,
		Price: input.Price,
		Sold: input.Sold,
	}
	if err := r.GetDB().Model(&item).Updates(data).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *itemRepository) DeleteItem(id string) error {
	var item models.Item
	if err := r.GetDB().Where("id = ? ", id).Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
