package models

type CreateItemInput struct {
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

type UpdateItemInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Sold  bool    `json:"sold"`
}
