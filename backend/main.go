package main

import (
	model "apicha/foodshop/Model"
	"apicha/foodshop/handler"

	"github.com/gin-contrib/cors"
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
	db.AutoMigrate(&model.Tablemodel{}, &model.Menumodel{}, &model.Customerordermodel{}, &model.Order{}, &model.Auth{})

	r := gin.Default()

	tableHandler := handler.NewTablehandler(db)
	menuHandler := handler.Newmenuhandler(db)
	customer_orders := handler.Neworderhandle(db)
	authentication := handler.NewAuthhandler(db)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Cookie", "Access-Control-Allow-Origin"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.Static("/Picture", "./Picture/")
	admin := r.Group("/admin")
	admin.Use(authentication.Middleware())
	{

		admin.GET("/tables", tableHandler.GetTable)
		admin.POST("/tables", tableHandler.CreateTable)
		admin.PUT("/tables/whereidupdate/:id", tableHandler.UpdateTable)
		admin.DELETE("/tables/:id", tableHandler.DeleteTable)

		admin.GET("/menus", menuHandler.GetAll)
		admin.POST("/menus", menuHandler.Create)
		admin.PUT("/menus/:id", menuHandler.Update)
		admin.DELETE("/menus/:id", menuHandler.Deletemenu)

		admin.GET("/customer_orders", customer_orders.GetAll)
		admin.GET("/customer_orders/GetMoneyByDay", customer_orders.GetMoneyByDay)
		admin.POST("/customer_orders", customer_orders.Create)
		admin.PUT("/customer_orders/:id", customer_orders.Update)
		admin.PUT("/customer_orders/status/:id", customer_orders.Donestatus)
		admin.DELETE("/customer_orders/:id", customer_orders.Delete)
	}

	r.POST("/register", authentication.Register)
	r.POST("/login", authentication.Login)
	r.Run()
}
