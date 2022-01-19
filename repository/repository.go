package repository

// The repository layer is responsible for connecting directly to the database to retrieve and/or modify records.

import (
	"errors"
	"log"
	"strconv"

	"github.com/nsavelyeva/go-shopping/models"
	"gorm.io/gorm"
)

// ItemRepository is an interface for the struct itemRepository
type ItemRepository interface {
	ListItems() ([]models.Item, error)
	FindItem(id string) (*models.Item, bool, error)
	CreateItem(input *models.CreateItemInput) (*models.Item, error)
	UpdateItem(id string, input *models.UpdateItemInput) (*models.Item, error)
	DeleteItem(id string) error
}

type itemRepository struct {
	db *gorm.DB
}

func connectDB(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, config)
	return db, err
}

// NewItemRepository is a constructor for ItemRepository
func NewItemRepository(dialector gorm.Dialector, config *gorm.Config) *ItemRepository {
	db, err := connectDB(dialector, config)
	if err != nil {
		log.Fatalf("Failed to connect to the database due to error: %s", err)
		return nil
	}

	var r ItemRepository = &itemRepository{db: db}
	return &r
}

func (r *itemRepository) isItemComplete(item *models.Item) bool {
	return item.Name != nil && item.Price != nil && item.Sold != nil // i.e. all non-GORM fields are not nil
}

func (r *itemRepository) ListItems() ([]models.Item, error) {
	var items []models.Item
	r.db.Find(&items)

	return items, nil
}

func (r *itemRepository) FindItem(id string) (*models.Item, bool, error) {
	var item models.Item
	itemID, _ := strconv.Atoi(id)

	err := r.db.First(&item, itemID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	if !r.isItemComplete(&item) {
		return nil, false, errors.New("broken item found")
	}

	return &item, true, nil
}

func (r *itemRepository) CreateItem(input *models.CreateItemInput) (*models.Item, error) {
	f := false
	item := models.Item{
		Name:  &input.Name,
		Price: &input.Price,
		Sold:  &f,
	}
	if err := r.db.Save(&item).Error; err != nil {
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
		Name:  &input.Name,
		Price: &input.Price,
		Sold:  &input.Sold,
	}
	if err := r.db.Model(&item).Updates(data).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *itemRepository) DeleteItem(id string) error {
	var item models.Item
	if err := r.db.Where("id = ? ", id).Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
