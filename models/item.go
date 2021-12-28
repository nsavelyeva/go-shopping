package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Sold  bool    `json:"sold"`
}
