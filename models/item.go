package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name    *string  `json:"name"`
	Price   *float32 `json:"price"`
	Sold    *bool    `json:"sold"`
}

type ItemResponse struct {
	Data   Item    `json:"data"`
}

type ItemsResponse struct {
	Data  []Item    `json:"data"`
}
