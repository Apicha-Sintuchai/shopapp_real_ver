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
	Customer_at_table string  `json:"customer_at_table"`
	Orders            []Order `json:"orders" gorm:"foreignKey:CustomerordermodelID"`
	Donestatus        string  `json:"donestatus"`
	Status            string  `json:"status"`
	Totalprice        string  `json:"totalprice"`
}

type Order struct {
	gorm.Model
	CustomerordermodelID uint   `json:"customerordermodel_id"`
	NameMenu             string `json:"name_menu"`
	Price                string `json:"price"`
	Option               string `json:"option"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
