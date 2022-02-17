package models

// CreateItemInput is a struct to keep JSON data for HTTP requests to create Item
type CreateItemInput struct {
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

// UpdateItemInput is a struct to keep JSON data for HTTP requests to update Item
type UpdateItemInput struct {
	Name  *string  `json:"name"`
	Price *float32 `json:"price"`
	Sold  *bool    `json:"sold"`
}
