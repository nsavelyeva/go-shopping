package models

import "gorm.io/gorm"

// Item is a struct to keep full data about Item and match to a database schema
type Item struct {
	gorm.Model
	Name  *string  `json:"name"`
	Price *float32 `json:"price"`
	Sold  *bool    `json:"sold"`
}

// ItemResponse is a struct to keep data about a single Item in an HTTP response
type ItemResponse struct {
	Data Item `json:"data"`
}

// ItemsResponse is a struct to keep data about multiple Items in an HTTP response
type ItemsResponse struct {
	Data []Item `json:"data"`
}
