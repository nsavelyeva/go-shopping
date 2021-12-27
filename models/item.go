package models

type Item struct {
	ID    uint32  `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Sold  bool    `json:"sold"`
}
