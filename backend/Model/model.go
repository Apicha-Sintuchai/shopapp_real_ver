package model

import "gorm.io/gorm"

type Tablemodel struct {
	gorm.Model
	TableNumber string `json:"table_number"`
}

type Menumodel struct {
	gorm.Model
	PictureMenu string `form:"picturemenu"`
	NameMenu    string `form:"namemenu"`
	Price       string `form:"price"`
	Status      string `form:"status"`
	Category    string `form:"category"`
}

type Customerordermodel struct {
	gorm.Model
	Customer_at_table string
	Orders            []Order
}

type Order struct {
	NameMenu   string
	Totalprice string
	Option     string
}
