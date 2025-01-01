package main

import (
	model "apicha/foodshop/Model"
	"apicha/foodshop/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=0126 dbname=qrcode_order port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Tablemodel{}, &model.Menumodel{})

	r := gin.Default()

	tableHandler := handler.NewTablehandler(db)
	menuHandler := handler.Newmenuhandler(db)
	r.GET("/tables", tableHandler.GetTable)
	r.POST("/tables", tableHandler.CreateTable)
	r.PUT("/tables/whereidupdate/:id", tableHandler.UpdateTable)
	r.DELETE("/tables/:id", tableHandler.DeleteTable)

	r.GET("/menus", menuHandler.GetAll)
	r.POST("/menus", menuHandler.Create)
	r.PUT("/menus/:id", menuHandler.Update)
	r.DELETE("/menus/:id", menuHandler.Deletemenu)
	r.Run()
}
