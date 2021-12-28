package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/nsavelyeva/go-shopping/models"
)

func ListItems(db *gorm.DB) ([]models.Item, error) {
	items := []models.Item{}
	query := db.Select("items.*").
		        Group("items.id").
				Find(&items)
	if err := query.Error; err != nil {
		return items, err
	}

	return items, nil
}

func FindItem(db *gorm.DB, id string) (models.Item, bool, error) {
	item := models.Item{}
	query := db.Select("items.*").
			    Group("items.id").
				Where("items.id = ?", id).
				First(&item)
	err := query.Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return item, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return item, false, nil
	}
	return item, true, nil
}

func CreateItem(db *gorm.DB, input *models.CreateItemInput) (*models.Item, error) {
	item := models.Item{Name: input.Name, Price: input.Price, Sold: false}
	if err := db.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func UpdateItem(db *gorm.DB, id string, input *models.UpdateItemInput) (*models.Item, error) {
	item, found, err := FindItem(db, id)
	if err != nil || !found {
		return nil, errors.New("Item not found")
	}
	data := models.Item{Name: input.Name, Price: input.Price, Sold: input.Sold}
	if err := db.Model(&item).Updates(data).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func DeleteItem(db *gorm.DB, id string) error {
	var item models.Item
	if err := db.Where("id = ? ", id).Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
